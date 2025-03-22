package server

import (
	"context"
	"os"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
	wallet "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/commands"
	"github.com/bitcoin-sv/spv-wallet-go-client/config"
	"github.com/bitcoin-sv/spv-wallet-go-client/queries"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

type AdminData struct {
	CommonData
	adminAPI *wallet.AdminAPI
}

// NewAdminAPI returns the accounts management interface for the  admin
func NewAdminAPI(ctx context.Context) (types.AccountWalletInfo, error) {
	adminAPIData, err := wallet.NewAdminAPIWithXPriv(
		config.New(config.WithAddr(os.Getenv("WALLET_URL"))), os.Getenv("USER_XPRIV"),
	)
	if err != nil {
		return nil, err
	}

	res, err := adminAPIData.CreateXPub(ctx, &commands.CreateUserXpub{XPub: os.Getenv("USER_XPUB")})
	if err != nil {
		return nil, err
	}

	return &AdminData{
		CommonData: CommonData{
			ctx:     ctx,
			accInfo: res,
		},
		adminAPI: adminAPIData,
	}, nil
}

// GetBalance return the current unspent account balance for the admin
func (a *AdminData) GetBalance() uint64 {
	return a.accInfo.CurrentBalance
}

// GetTransactions returns the list of transactions associated with the admin account
func (a *AdminData) GetTransactions() ([]*response.Transaction, error) {
	page, err := a.adminAPI.Transactions(a.ctx,
		queries.QueryWithPageFilter[filter.AdminTransactionFilter](
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
func (a *AdminData) GetTransaction(txID string) (*response.Transaction, error) {
	return a.adminAPI.Transaction(a.ctx, txID)
}
