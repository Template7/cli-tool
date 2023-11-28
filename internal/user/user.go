package user

import (
	"context"
	"github.com/Template7/common/logger"
)

type User struct {
	token string
	log   *logger.Logger
}

func New(ctx context.Context, username string, password string) *User {
	log := logger.New().With("username", username)

	u := &User{
		log: log,
	}

	if !u.login(ctx, username, password) {
		log.Error("fail to login")
		return nil
	}

	return u
}
