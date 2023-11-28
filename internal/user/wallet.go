package user

import (
	"context"
	v1 "github.com/Template7/protobuf/gen/proto/template7/wallet"
)

func (u *User) GetWallet(ctx context.Context) {
	log := u.log.WithContext(ctx)
	log.Debug("get wallet")

	for _, uw := range u.be.GetUserWallets(ctx, u.token) {
		u.wallets[uw.Id] = map[v1.Currency]int{}
	}

}
