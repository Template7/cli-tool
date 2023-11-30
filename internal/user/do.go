package user

import (
	"context"
	v1 "github.com/Template7/protobuf/gen/proto/template7/wallet"
	"math/rand"
	"time"
)

const (
	doDeposit = iota
	doWithdraw
	doTransfer
)

var (
	currencies = []string{v1.Currency_ntd.String(), v1.Currency_usd.String(), v1.Currency_jpy.String(), v1.Currency_cny.String()}
)

func (u *User) Do(ctx context.Context, threshold int, toWallets []string, rest int) {
	log := u.log.WithContext(ctx).With("threshold", threshold)
	log.Debug("user do")

	for {
		select {
		case <-ctx.Done():
			log.Info("context done")
			u.GetWallet(ctx)
			log.With("walletInfo", u.wallets).Info("show wallets")
			return
		default:
			act, cur, blc := u.checkAction(ctx, threshold)
			switch act {
			case doDeposit:
				if err := u.Deposit(ctx, cur, 1+rand.Intn(1000)); err != nil {
					log.WithError(err).Error("fail to deposit")
				}
			case doWithdraw:
				if err := u.Withdraw(ctx, cur, 1+rand.Intn(threshold)); err != nil {
					log.WithError(err).Error("fail to deposit")
				}
			case doTransfer:
				if err := u.Transfer(ctx, toWallets[rand.Intn(len(toWallets))], cur, 1+rand.Intn(blc)); err != nil {
					log.WithError(err).Error("fail to deposit")
				}
			}

			u.GetWallet(ctx)
			time.Sleep(time.Duration(rest) * time.Millisecond)
		}
	}
}

func (u *User) checkAction(ctx context.Context, threshold int) (int, string, int) {
	log := u.log.WithContext(ctx).With("threshold", threshold)
	log.Debug("check action")

	for _, bls := range u.wallets {
		for cur, bl := range bls {
			if bl > threshold {
				return doWithdraw, cur, bl
			}
			if bl > 0 {
				return doTransfer, cur, bl
			}
			return doDeposit, cur, bl
		}
	}
	return doDeposit, currencies[rand.Intn(len(currencies))], 0
}
