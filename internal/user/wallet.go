package user

import (
	"cli-tool/internal/backend"
	"github.com/Template7/backend/pkg/apiBody"
)

func (u *User) Transfer(req apiBody.TransactionReq) (err error) {
	log.Debug("user make transfer: ", req.String())

	_, err = backend.New().Transaction(req, u.Token.AccessToken)
	return
}

func (u *User) Deposit(req apiBody.DepositReq) (err error) {
	log.Debug("deposit money")

	err = backend.New().Deposit(req, u.Token.AccessToken)
	return
}
