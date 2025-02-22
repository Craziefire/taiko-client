package producer

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

func TestNewZkevmRpcdProducer(t *testing.T) {
	dummyZkevmRpcdProducer, err := NewZkevmRpcdProducer(
		"http://localhost:18545",
		"",
		"",
		"",
		false,
		&bindings.TaikoDataConfig{},
	)
	require.Nil(t, err)
	require.False(t, dummyZkevmRpcdProducer.Cancellable())

	dummyZkevmRpcdProducer.CustomProofHook = func() ([]byte, uint64, error) {
		return []byte{0}, CircuitsIdx, nil
	}

	resCh := make(chan *ProofWithHeader, 1)

	blockID := common.Big32
	header := &types.Header{
		ParentHash:  randHash(),
		UncleHash:   randHash(),
		Coinbase:    common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec"),
		Root:        randHash(),
		TxHash:      randHash(),
		ReceiptHash: randHash(),
		Difficulty:  common.Big0,
		Number:      common.Big256,
		GasLimit:    1024,
		GasUsed:     1024,
		Time:        uint64(time.Now().Unix()),
		Extra:       randHash().Bytes(),
		MixDigest:   randHash(),
		Nonce:       types.BlockNonce{},
	}

	require.Nil(t, dummyZkevmRpcdProducer.RequestProof(
		context.Background(),
		&ProofRequestOptions{},
		blockID,
		&bindings.TaikoDataBlockMetadata{},
		header,
		resCh,
	))

	res := <-resCh
	require.Equal(t, res.BlockID, blockID)
	require.Equal(t, res.Header, header)
	require.NotEmpty(t, res.Proof)

	require.Nil(t, dummyZkevmRpcdProducer.Cancel(context.Background(), common.Big1))
}

func TestZkevmRpcdProducerCalls(t *testing.T) {
	dummyZkevmRpcdProducer, err := NewZkevmRpcdProducer(
		"",
		"",
		"",
		"",
		false,
		&bindings.TaikoDataConfig{
			BlockMaxGasLimit:    uint32(randHash().Big().Uint64()),
			BlockMaxTxListBytes: randHash().Big(),
		},
	)
	require.Nil(t, err)
	require.False(t, dummyZkevmRpcdProducer.Cancellable())

	dummyZkevmRpcdProducer.CustomProofHook = func() ([]byte, uint64, error) {
		return []byte{0}, CircuitsIdx, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, _, err = dummyZkevmRpcdProducer.callProverDaemon(
		ctx,
		&ProofRequestOptions{BlockID: common.Big32},
		&bindings.TaikoDataBlockMetadata{},
	)

	require.Nil(t, err)
}
