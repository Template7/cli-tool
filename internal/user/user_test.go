package user

import (
	"context"
	v1 "github.com/Template7/protobuf/gen/proto/template7/wallet"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"testing"
)

func TestNew(t *testing.T) {
	viper.AddConfigPath("../../config")
	ctx := context.WithValue(context.Background(), "traceId", uuid.NewString())

	u1 := New(ctx, "fakeUser001", "fakeUser001")
	t.Log(u1.String())

	if err := u1.Deposit(ctx, v1.Currency_usd.String(), 100); err != nil {
		t.Error(err)
		return
	}

	u1.GetWallet(ctx)
	t.Log(u1.String())

	if err := u1.Withdraw(ctx, v1.Currency_usd.String(), 50); err != nil {
		t.Error(err)
		return
	}

	u1.GetWallet(ctx)
	t.Log(u1.String())

	u2 := New(ctx, "fakeUser002", "fakeUser002")
	t.Log(u2.String())

	for wId, _ := range u2.wallets {
		if err := u1.Transfer(ctx, wId, v1.Currency_usd.String(), 30); err != nil {
			t.Error(err)
			return
		}
		break
	}

	u1.GetWallet(ctx)
	u2.GetWallet(ctx)

	t.Log(u1.String())
	t.Log(u2.String())
}
