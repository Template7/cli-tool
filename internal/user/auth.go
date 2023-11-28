package user

import (
	"cli-tool/internal/backend"
	"context"
)

func (u *User) login(ctx context.Context, username string, password string) bool {
	log := u.log.WithContext(ctx)
	log.Debug("user login")

	token := backend.New().NativeLogin(ctx, username, password)
	if token == "" {
		log.Warn("empty token")
		return false
	}
	u.token = token
	return true
}
