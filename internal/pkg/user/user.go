package user

import (
	"cli-tool/internal/pkg/t7Redis"
	"fmt"
	"github.com/Template7/common/logger"
	"github.com/Template7/common/structs"
	"github.com/dgrijalva/jwt-go"
)

var (
	log = logger.GetLogger()
)

type User struct {
	Token  structs.Token
	Data   structs.User
	Wallet structs.WalletData
}

func (u *User) updateToken(token structs.Token) (err error) {
	log.Debug("update user token")

	// parse jwt without validate
	tk, _ := jwt.ParseWithClaims(token.AccessToken, &structs.UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	utc, ok := tk.Claims.(*structs.UserTokenClaims)
	if !ok {
		log.Error("fail to assert user token claims")
		return fmt.Errorf("fail to assert user token claims")
	}

	u.Token = token
	u.Data.UserId = utc.UserId

	log.Debug("finish update user token")
	return
}

func getVerifyCode(mobile string) (verifyCode string) {
	log.Debug("get verify code for mobile: ", mobile)
	verifyCode = t7Redis.New().Get(fmt.Sprintf("%s:%s", t7Redis.VerifyCodePrefix, mobile)).Val()
	log.Debug("got verify code: ", verifyCode)
	return
}
