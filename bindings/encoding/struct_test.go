package encoding

import (
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/taikoxyz/taiko-client/common/utils"
)

var (
	testHeader = &types.Header{
		ParentHash:  randomHash(),
		UncleHash:   types.EmptyUncleHash,
		Coinbase:    common.BytesToAddress(randomHash().Bytes()),
		Root:        randomHash(),
		TxHash:      randomHash(),
		ReceiptHash: randomHash(),
		Bloom:       types.BytesToBloom(randomHash().Bytes()),
		Difficulty:  new(big.Int).SetUint64(utils.RandUint64(nil)),
		Number:      new(big.Int).SetUint64(utils.RandUint64(nil)),
		GasLimit:    utils.RandUint64(nil),
		GasUsed:     utils.RandUint64(nil),
		Time:        uint64(time.Now().Unix()),
		Extra:       randomHash().Bytes(),
		MixDigest:   randomHash(),
		Nonce:       types.EncodeNonce(utils.RandUint64(nil)),
		BaseFee:     new(big.Int).SetUint64(utils.RandUint64(nil)),
	}
)

func TestFromGethHeader(t *testing.T) {
	header := FromGethHeader(testHeader)

	require.Equal(t, testHeader.ParentHash, common.BytesToHash(header.ParentHash[:]))
	require.Equal(t, testHeader.UncleHash, common.BytesToHash(header.OmmersHash[:]))
	require.Equal(t, testHeader.Coinbase, header.Beneficiary)
	require.Equal(t, testHeader.Root, common.BytesToHash(header.StateRoot[:]))
	require.Equal(t, testHeader.TxHash, common.BytesToHash(header.TransactionsRoot[:]))
	require.Equal(t, testHeader.ReceiptHash, common.BytesToHash(header.ReceiptsRoot[:]))
	require.Equal(t, BloomToBytes(testHeader.Bloom), header.LogsBloom)
	require.Equal(t, testHeader.Difficulty, header.Difficulty)
	require.Equal(t, testHeader.Number, header.Height)
	require.Equal(t, testHeader.GasLimit, header.GasLimit)
	require.Equal(t, testHeader.GasUsed, header.GasUsed)
	require.Equal(t, testHeader.Time, header.Timestamp)
	require.Equal(t, testHeader.Extra, header.ExtraData)
	require.Equal(t, testHeader.MixDigest, common.BytesToHash(header.MixHash[:]))
	require.Equal(t, testHeader.Nonce.Uint64(), header.Nonce)
	require.Equal(t, testHeader.BaseFee.Uint64(), header.BaseFeePerGas.Uint64())
}

func TestFromToGethHeaderLegacyTx(t *testing.T) {
	testHeader := testHeader // Copy the original struct
	testHeader.BaseFee = nil
	header := FromGethHeader(testHeader)

	require.Equal(t, testHeader.ParentHash, common.BytesToHash(header.ParentHash[:]))
	require.Equal(t, testHeader.UncleHash, common.BytesToHash(header.OmmersHash[:]))
	require.Equal(t, testHeader.Coinbase, header.Beneficiary)
	require.Equal(t, testHeader.Root, common.BytesToHash(header.StateRoot[:]))
	require.Equal(t, testHeader.TxHash, common.BytesToHash(header.TransactionsRoot[:]))
	require.Equal(t, testHeader.ReceiptHash, common.BytesToHash(header.ReceiptsRoot[:]))
	require.Equal(t, BloomToBytes(testHeader.Bloom), header.LogsBloom)
	require.Equal(t, testHeader.Difficulty, header.Difficulty)
	require.Equal(t, testHeader.Number, header.Height)
	require.Equal(t, testHeader.GasLimit, header.GasLimit)
	require.Equal(t, testHeader.GasUsed, header.GasUsed)
	require.Equal(t, testHeader.Time, header.Timestamp)
	require.Equal(t, testHeader.Extra, header.ExtraData)
	require.Equal(t, testHeader.MixDigest, common.BytesToHash(header.MixHash[:]))
	require.Equal(t, testHeader.Nonce.Uint64(), header.Nonce)
	require.Equal(t, new(big.Int).SetInt64(0).Uint64(), header.BaseFeePerGas.Uint64())

	gethHeader := ToGethHeader(header)
	require.Equal(t, testHeader, gethHeader)
}

func TestToExecutableData(t *testing.T) {
	data := ToExecutableData(testHeader)
	require.Equal(t, testHeader.ParentHash, data.ParentHash)
	require.Equal(t, testHeader.Coinbase, data.FeeRecipient)
	require.Equal(t, testHeader.Root, data.StateRoot)
	require.Equal(t, testHeader.ReceiptHash, data.ReceiptsRoot)
	require.Equal(t, testHeader.Bloom.Bytes(), data.LogsBloom)
	require.Equal(t, testHeader.MixDigest, data.Random)
	require.Equal(t, testHeader.Number.Uint64(), data.Number)
	require.Equal(t, testHeader.GasLimit, data.GasLimit)
	require.Equal(t, testHeader.GasUsed, data.GasUsed)
	require.Equal(t, testHeader.Time, data.Timestamp)
	require.Equal(t, testHeader.Extra, data.ExtraData)
	require.Equal(t, testHeader.BaseFee, data.BaseFeePerGas)
	require.Equal(t, testHeader.Hash(), data.BlockHash)
	require.Equal(t, testHeader.TxHash, data.TxHash)
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}

// randomBytes generates a random bytes.
func randomBytes(size int) (b []byte) {
	b = make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		log.Crit("Generate random bytes error", "error", err)
	}
	return
}
