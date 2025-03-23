package server

import (
	"context"
	"fmt"
	"os"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
	wallet "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/config"
	"github.com/bitcoin-sv/spv-wallet-go-client/queries"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

type CommonData struct {
	ctx     context.Context
	accInfo *response.Xpub
}

type UserData struct {
	CommonData
	userAPIData *wallet.UserAPI
}

// NewUserAPI returns the accounts management interface for the  admin
func NewUserAPI(ctx context.Context) (types.AccountWalletInfo, error) {
	fmt.Println("Running the User API with the privated Key!")

	userAPI, err := wallet.NewUserAPIWithXPriv(
		config.New(config.WithAddr(os.Getenv("WALLET_URL"))), os.Getenv("USER_XPRIV"),
	)
	if err != nil {
		return nil, err
	}

	res, err := userAPI.XPub(ctx)
	if err != nil {
		return nil, err
	}

	return &UserData{
		CommonData: CommonData{
			ctx:     ctx,
			accInfo: res,
		},
		userAPIData: userAPI,
	}, nil
}

// GetBalance return the current unspent account balance for the admin
func (a *UserData) GetBalance() uint64 {
	return a.accInfo.CurrentBalance
}

// GetTransactions returns the list of transactions associated with the admin account
func (a *UserData) GetTransactions() ([]*response.Transaction, error) {
	page, err := a.userAPIData.Transactions(a.ctx,
		queries.QueryWithPageFilter[filter.TransactionFilter](
			filter.Page{
				Size: 10,
				Sort: "asc",
			}))
	if err != nil {
		return nil, err
	}
	return page.Content, err
}

// GetTransaction returns the transactions associated with a given tx ID
func (a *UserData) GetTransaction(txID string) (*response.Transaction, error) {
	return a.userAPIData.Transaction(a.ctx, txID)
}
