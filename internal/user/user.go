package user

import (
	"cli-tool/internal/backend"
	"context"
	"encoding/json"
	"github.com/Template7/backend/api/types"
	"github.com/Template7/common/logger"
)

type User struct {
	name    string
	token   string
	wallets map[string]map[string]int // [walletId][currency]amount
	records map[string]map[string][]types.HttpGetWalletBalanceRecordRespData

	be  *backend.Client
	log *logger.Logger
}

func (u *User) String() string {
	d := struct {
		Username       string                                                           `json:"username"`
		WalletBalances map[string]map[string]int                                        `json:"walletBalances"`
		WalletRecords  map[string]map[string][]types.HttpGetWalletBalanceRecordRespData `json:"walletRecords"`
	}{
		Username:       u.name,
		WalletBalances: u.wallets,
		WalletRecords:  u.records,
	}
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}

func New(ctx context.Context, username string, password string) *User {
	log := logger.New().With("username", username)

	u := &User{
		name:    username,
		be:      backend.New(),
		wallets: map[string]map[string]int{},
		records: map[string]map[string][]types.HttpGetWalletBalanceRecordRespData{},
		log:     log,
	}

	if !u.login(ctx, username, password) {
		log.Panic("fail to login")
		return nil
	}
	u.GetWallet(ctx)

	return u
}
