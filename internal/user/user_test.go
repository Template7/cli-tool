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

	u := New(ctx, "fakeUser001", "fakeUser001")
	u.GetWallet(ctx)
	t.Log(u.String())

	if err := u.Deposit(ctx, v1.Currency_usd.String(), 100); err != nil {
		t.Error(err)
		return
	}

	u.GetWallet(ctx)
	t.Log(u.String())

	if err := u.Withdraw(ctx, v1.Currency_usd.String(), 50); err != nil {
		t.Error(err)
		return
	}

	u.GetWallet(ctx)
	t.Log(u.String())
}
