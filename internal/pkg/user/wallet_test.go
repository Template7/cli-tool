package user

import (
	fakeDataGenerator "cli-tool/internal/pkg/fakeData"
	"github.com/Template7/backend/pkg/apiBody"
	"github.com/Template7/common/structs"
	"github.com/spf13/viper"
	"testing"
)

func TestUser_Transfer(t *testing.T) {
	viper.AddConfigPath("../../../configs")

	// sign up sender and receiver
	senderData := fakeDataGenerator.RandomUser()
	fakeSender := User{
		Data: senderData,
	}
	if err := fakeSender.SignUp(); err != nil {
		t.Error("fail to sign up user: ", fakeSender.Data.UserId)
		return
	}
	if err := fakeSender.UpdateInfo(senderData.BasicInfo); err != nil {
		t.Error("fail to update user info: ", err.Error())
		return
	}
	if err := fakeSender.GetInfo(); err != nil {
		t.Error(err)
		return
	}
	receiverData := fakeDataGenerator.RandomUser()
	fakeReceiver := User{
		Data: receiverData,
	}
	if err := fakeReceiver.SignUp(); err != nil {
		t.Error("fail to sign up user: ", fakeReceiver.Data.UserId)
		return
	}
	if err := fakeReceiver.UpdateInfo(receiverData.BasicInfo); err != nil {
		t.Error("fail to update user info: ", err.Error())
		return
	}
	if err := fakeReceiver.GetInfo(); err != nil {
		t.Error(err)
		return
	}

	// sender deposit
	depositData := apiBody.DepositReq{
		Source:   "fakeSource",
		Note:     "test",
		WalletId: fakeSender.Wallet.Id,
		Money: structs.Money{
			Currency: structs.CurrencyNTD,
			Amount:   100,
			Unit:     structs.UnitOne,
		},
	}
	if err := fakeSender.Deposit(depositData); err != nil {
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
