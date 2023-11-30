package user

import (
	"context"
	"fmt"
	"strconv"
)

func (u *User) GetWallet(ctx context.Context) map[string]map[string]int {
	log := u.log.WithContext(ctx)
	log.Debug("get wallet")

	for _, uw := range u.be.GetUserWallets(ctx, u.token) {
		if _, ok := u.wallets[uw.Id]; !ok {
			u.wallets[uw.Id] = map[string]int{}
		}
		for _, blc := range uw.Balances {
			am, err := strconv.Atoi(blc.Amount)
			if err != nil {
				log.WithError(err).With("amount", blc.Amount).Error("fail to convert amount")
				continue
			}
			u.wallets[uw.Id][blc.Currency] = am
		}
	}
	return u.wallets
}

func (u *User) Deposit(ctx context.Context, currency string, amount int) error {
	log := u.log.WithContext(ctx).With("currency", currency).With("amount", amount)
	log.Debug("user deposit")

	if len(u.wallets) == 0 {
		log.Warn("user has no wallet")
		return fmt.Errorf("user has no wallet")
	}

	if len(u.wallets) > 1 {
		log.Warn("user has multiple wallets, deposit the 1st as default")
	}

	for wId, _ := range u.wallets {
		if err := u.be.Deposit(ctx, wId, currency, amount, u.token); err != nil {
			log.WithError(err).Error("fail to deposit")
			return err
		}
		return nil
	}
	return nil
}

func (u *User) Withdraw(ctx context.Context, currency string, amount int) error {
	log := u.log.WithContext(ctx).With("currency", currency).With("amount", amount)
	log.Debug("user withdraw")

	if len(u.wallets) == 0 {
		log.Warn("user has no wallet")
		return fmt.Errorf("user has no wallet")
	}

	if len(u.wallets) > 1 {
		log.Warn("user has multiple wallets, deposit the 1st as default")
	}

	for wId, _ := range u.wallets {
		if err := u.be.Withdraw(ctx, wId, currency, amount, u.token); err != nil {
			log.WithError(err).Error("fail to withdraw")
			return err
		}
		return nil
	}
	return nil
}

func (u *User) Transfer(ctx context.Context, toWalletId string, currency string, amount int) error {
	log := u.log.WithContext(ctx).With("currency", currency).With("amount", amount)
	log.Debug("user transfer money")

	if len(u.wallets) == 0 {
		log.Warn("user has no wallet")
		return fmt.Errorf("user has no wallet")
	}

	if len(u.wallets) > 1 {
		log.Warn("user has multiple wallets, deposit the 1st as default")
	}

	for wId, _ := range u.wallets {
		if err := u.be.Transfer(ctx, wId, toWalletId, currency, amount, u.token); err != nil {
			log.WithError(err).Error("fail to transfer")
			return err
		}
		return nil
	}
	return nil
}
