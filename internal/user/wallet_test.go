package user

import (
	fakeDataGenerator "cli-tool/internal/fakeData"
	"github.com/Template7/backend/pkg/apiBody"
	"github.com/Template7/common/structs"
	"github.com/spf13/viper"
	"testing"
)

var (
	senderData = fakeDataGenerator.RandomUser()
	fakeSender = User{
		Data: senderData,
	}
	receiverData = fakeDataGenerator.RandomUser()
	fakeReceiver = User{
		Data: receiverData,
	}
)

func signUpFakeUser() error {
	if err := fakeSender.SignUp(); err != nil {
		return err
	}
	if err := fakeSender.UpdateInfo(senderData.BasicInfo); err != nil {
		return err
	}
	if err := fakeSender.GetInfo(); err != nil {
		return err
	}
	if err := fakeReceiver.SignUp(); err != nil {
		return err
	}
	if err := fakeReceiver.UpdateInfo(receiverData.BasicInfo); err != nil {
		return err
	}
	if err := fakeReceiver.GetInfo(); err != nil {
		return err
	}
	return nil
}

func deposit(user User) error {
	depositData := apiBody.DepositReq{
		Source:   "fakeSource",
		Note:     "test",
		WalletId: user.Wallet.Id,
		Money: structs.Money{
			Currency: structs.CurrencyNTD,
			Amount:   100,
			Unit:     structs.UnitOne,
		},
	}
	if err := user.Deposit(depositData); err != nil {
		return err
	}
	return nil
}

func TestUser_Transfer(t *testing.T) {
	viper.AddConfigPath("../../../configs")

	// sign up sender and receiver
	if err := signUpFakeUser(); err != nil {
		t.Error(err)
		return
	}

	// sender deposit
	if err := deposit(fakeSender); err != nil {
		t.Error(err)
		return
	}

	transferData := apiBody.TransactionReq{
		FromWalletId: fakeSender.Wallet.Id,
		ToWalletId:   fakeReceiver.Wallet.Id,
		Money: structs.Money{
			Currency: structs.CurrencyNTD,
			Amount:   100,
			Unit:     structs.UnitOne,
		},
	}
	if err := fakeSender.Transfer(transferData); err != nil {
		t.Error(err)
	}
	return
}

func BenchmarkTransfer(b *testing.B) {
	viper.AddConfigPath("../../../configs")

	if err := signUpFakeUser(); err != nil {
		b.Error(err)
		return
	}
	if err := deposit(fakeSender); err != nil {
		b.Error(err)
		return
	}

	transferDataFromSender := apiBody.TransactionReq{
		FromWalletId: fakeSender.Wallet.Id,
		ToWalletId:   fakeReceiver.Wallet.Id,
		Money: structs.Money{
			Currency: structs.CurrencyNTD,
			Amount:   100,
			Unit:     structs.UnitOne,
		},
	}
	transferDataFromReceiver := apiBody.TransactionReq{
		FromWalletId: fakeReceiver.Wallet.Id,
		ToWalletId:   fakeSender.Wallet.Id,
		Money: structs.Money{
			Currency: structs.CurrencyNTD,
			Amount:   100,
			Unit:     structs.UnitOne,
		},
	}
	for i := 0; i < b.N; i++ {
		if err := fakeSender.Transfer(transferDataFromSender); err != nil {
			b.Error(err)
			return
		}
		if err := fakeReceiver.Transfer(transferDataFromReceiver); err != nil {
			b.Error(err)
			return
		}
	}
}
