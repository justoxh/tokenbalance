package tokenbalance

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Config struct {
	GethLocation string
	UsePort      int
	UseIP        string
	Logs         bool
}

type TokenBalance struct {
	Contract common.Address
	Wallet   common.Address
	Name     string
	Symbol   string
	Balance  *big.Int
	ETH      *big.Int
	Decimals int64
	Block    int64
	ctx      context.Context
}

type TokenBalanceJson struct {
	Contract string `json:"token,omitempty"`
	Wallet   string `json:"wallet,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Balance  string `json:"balance"`
	ETH      string `json:"eth_balance,omitempty"`
	Decimals int64  `json:"decimals,omitempty"`
	Block    int64  `json:"block,omitempty"`
}
