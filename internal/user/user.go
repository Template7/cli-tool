package user

import (
	"cli-tool/internal/backend"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/common/logger"
)

type User struct {
	name    string
	token   string
	wallets map[string]map[string]int // [walletId][currency]amount

	be  *backend.Client
	log *logger.Logger
}

func (u *User) String() string {
	b, _ := json.MarshalIndent(u.wallets, "", "  ")
	return fmt.Sprintf("username: %s, wallets: %s", u.name, string(b))
}

func New(ctx context.Context, username string, password string) *User {
	log := logger.New().With("username", username)

	u := &User{
		name:    username,
		be:      backend.New(),
		wallets: map[string]map[string]int{},
		log:     log,
	}

	if !u.login(ctx, username, password) {
		log.Error("fail to login")
		return nil
	}
	u.GetWallet(ctx)

	return u
}
