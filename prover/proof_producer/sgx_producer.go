package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
)

// SGXProofProducer generates a SGX proof for the given block.
type SGXProofProducer struct {
	RaikoHostEndpoint string // a proverd RPC endpoint
	L1Endpoint        string // a L1 node RPC endpoint
	L2Endpoint        string // a L2 execution engine's RPC endpoint
	*DummyProofProducer
}

// SGXRequestProofBody represents the JSON body for requesting the proof.
type SGXRequestProofBody struct {
	JsonRPC string                      `json:"jsonrpc"` //nolint:revive,stylecheck
	ID      *big.Int                    `json:"id"`
	Method  string                      `json:"method"`
	Params  []*SGXRequestProofBodyParam `json:"params"`
}

// SGXRequestProofBodyParam represents the JSON body of RequestProofBody's `param` field.
type SGXRequestProofBodyParam struct {
	Type     string   `json:"type"`
	Block    *big.Int `json:"block"`
	L2RPC    string   `json:"l2Rpc"`
	L1RPC    string   `json:"l1Rpc"`
	Prover   string   `json:"prover"`
	Graffiti string   `json:"graffiti"`
}

// SGXRequestProofBodyResponse represents the JSON body of the response of the proof requests.
type SGXRequestProofBodyResponse struct {
	JsonRPC string           `json:"jsonrpc"` //nolint:revive,stylecheck
	ID      *big.Int         `json:"id"`
	Result  *RaikoHostOutput `json:"result"`
	Error   *struct {
		Code    *big.Int `json:"code"`
		Message string   `json:"message"`
	} `json:"error,omitempty"`
}

// RaikoHostOutput represents the JSON body of SGXRequestProofBodyResponse's `result` field.
type RaikoHostOutput struct {
	Type  string `json:"type"`
	Proof string `json:"proof"`
}

// NewSGXProducer creates a new `SGXProofProducer` instance.
func NewSGXProducer(
	raikoHostEndpoint string,
	l1Endpoint string,
	l2Endpoint string,
) (*SGXProofProducer, error) {
	return &SGXProofProducer{
		RaikoHostEndpoint: raikoHostEndpoint,
		L1Endpoint:        l1Endpoint,
		L2Endpoint:        l2Endpoint,
	}, nil
}

// RequestProof implements the ProofProducer interface.
func (s *SGXProofProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request proof from raiko-host service",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	if s.DummyProofProducer != nil {
		return s.DummyProofProducer.RequestProof(ctx, opts, blockID, meta, header, s.Tier(), resultCh)
	}

	proof, err := s.callProverDaemon(ctx, opts)
	if err != nil {
		return err
	}

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Header:  header,
		Meta:    meta,
		Proof:   proof,
		Degree:  0,
		Opts:    opts,
		Tier:    s.Tier(),
	}

	metrics.ProverSgxProofGeneratedCounter.Inc(1)

	return nil
}

// callProverDaemon keeps polling the proverd service to get the requested proof.
func (s *SGXProofProducer) callProverDaemon(ctx context.Context, opts *ProofRequestOptions) ([]byte, error) {
	var (
		proof []byte
		start = time.Now()
	)
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}
		output, err := s.requestProof(opts)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.BlockID, "err", err, "endpoint", s.RaikoHostEndpoint)
			return err
		}

		if output == nil {
			log.Info(
				"Proof generating",
				"height", opts.BlockID,
				"time", time.Since(start),
				"producer", "SGXProofProducer",
			)
			return errProofGenerating
		}

		log.Debug("Proof generation output", "output", output)

		proof = common.Hex2Bytes(output.Proof[2:])
		log.Info(
			"Proof generated",
			"height", opts.BlockID,
			"time", time.Since(start),
			"producer", "SGXProofProducer",
		)
		return nil
	}, backoff.NewConstantBackOff(proofPollingInterval)); err != nil {
		return nil, err
	}

	return proof, nil
}

// requestProof sends a RPC request to proverd to try to get the requested proof.
func (s *SGXProofProducer) requestProof(opts *ProofRequestOptions) (*RaikoHostOutput, error) {
	reqBody := SGXRequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*SGXRequestProofBodyParam{{
			Type:     "Sgx",
			Block:    opts.BlockID,
			L2RPC:    s.L2Endpoint,
			L1RPC:    s.L1Endpoint,
			Prover:   opts.ProverAddress.Hex()[2:],
			Graffiti: opts.Graffiti,
		}},
	}

	jsonValue, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(s.RaikoHostEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request proof, id: %d, statusCode: %d", opts.BlockID, res.StatusCode)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output SGXRequestProofBodyResponse
	if err := json.Unmarshal(resBytes, &output); err != nil {
		return nil, err
	}

	if output.Error != nil {
		return nil, errors.New(output.Error.Message)
	}

	return output.Result, nil
}

// Tier implements the ProofProducer interface.
func (s *SGXProofProducer) Tier() uint16 {
	return encoding.TierSgxID
}

// Cancellable implements the ProofProducer interface.
func (s *SGXProofProducer) Cancellable() bool {
	return false
}

// Cancel cancels an existing proof generation.
//
//nolint:golint
func (s *SGXProofProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
