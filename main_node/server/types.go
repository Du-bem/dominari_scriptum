package server

import "github.com/bitcoin-sv/spv-wallet/models/response"

type AccountWalletInfo interface {
	GetBalance() uint64
	GetTransactions() ([]*response.Transaction, error)
	GetTransaction(txID string) (*response.Transaction, error)
}
