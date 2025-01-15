package client

import (
	"context"
	"encoding/json"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

type GetParsedTransactionConfig struct {
	Commitment                     rpc.Commitment
	MaxSupportedTransactionVersion *uint8
}

func (c GetParsedTransactionConfig) toRpc() rpc.GetTransactionConfig {
	if c.MaxSupportedTransactionVersion == nil {
		// default
		c.MaxSupportedTransactionVersion = pointer.Get[uint8](0)
	}
	return rpc.GetTransactionConfig{
		Commitment:                     c.Commitment,
		Encoding:                       rpc.TransactionEncodingJsonParsed,
		MaxSupportedTransactionVersion: c.MaxSupportedTransactionVersion,
	}
}

type ParsedTransaction struct {
	Slot        uint64                   `json:"slot"`
	Meta        *rpc.TransactionMeta     `json:"meta"`
	Transaction *types.ParsedTransaction `json:"transaction"`
	BlockTime   *int64                   `json:"blockTime"`
}

// GetParsedTransaction returns transaction details for a confirmed transaction
func (c *Client) GetParsedTransaction(ctx context.Context, txhash string) (*ParsedTransaction, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetTransaction], error) {
			return c.RpcClient.GetTransactionWithConfig(ctx, txhash, GetParsedTransactionConfig{}.toRpc())
		},
		convertParsedTransaction,
	)
}

// GetParsedTransaction returns transaction details for a confirmed transaction
func (c *Client) GetParsedTransactionWithConfig(ctx context.Context, txhash string, cfg GetParsedTransactionConfig) (*ParsedTransaction, error) {
	return process(
		func() (rpc.JsonRpcResponse[*rpc.GetTransaction], error) {
			return c.RpcClient.GetTransactionWithConfig(ctx, txhash, cfg.toRpc())
		},
		convertParsedTransaction,
	)
}

func convertParsedTransaction(v *rpc.GetTransaction) (*ParsedTransaction, error) {
	if v == nil {
		return nil, nil
	}

	tx := new(types.ParsedTransaction)
	marshal, err := json.Marshal(v.Transaction)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshal, tx)
	if err != nil {
		return nil, err
	}

	return &ParsedTransaction{
		Slot:        v.Slot,
		BlockTime:   v.BlockTime,
		Transaction: tx,
		Meta:        v.Meta,
	}, nil
}
