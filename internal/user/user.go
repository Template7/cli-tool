package user

import (
	"cli-tool/internal/backend"
	"context"
	"github.com/Template7/common/logger"
	v1 "github.com/Template7/protobuf/gen/proto/template7/wallet"
)

type User struct {
	token   string
	wallets map[string]map[v1.Currency]int

	be  *backend.Client
	log *logger.Logger
}

func New(ctx context.Context, username string, password string) *User {
	log := logger.New().With("username", username)

	u := &User{
		be:      backend.New(),
		wallets: map[string]map[v1.Currency]int{},
		log:     log,
	}

	if !u.login(ctx, username, password) {
		log.Error("fail to login")
		return nil
	}

	return u
}
