package user

import (
	"cli-tool/internal/pkg/backend"
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
	Token structs.Token
	Data  structs.User
}

func (u *User) SignUp() error {
	log.Debug("user sign up: ", u.Data.UserId)

	if err := backend.New().SendSms(u.Data.Mobile); err != nil {
		log.Error("fail to send sms: ", err.Error())
		return err
	}

	signUpData := backend.ConfirmSmsVerifyCodeReq{
		Mobile: u.Data.Mobile,
		Code:   getVerifyCode(u.Data.Mobile),
	}
	token, err := backend.New().MobileSignUp(signUpData)
	if err != nil {
		log.Error("fail to sign up: ", err.Error())
		return err
	}
	if err := u.updateToken(token); err != nil {
		log.Error("fail to update token: ", err.Error())
		return err
	}

	log.Debug("finish user sign up")
	return nil
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

func (u *User) SignIn() error {
	log.Debug("user sign in: ", u.Data.UserId)

	if err := backend.New().SendSms(u.Data.Mobile); err != nil {
		log.Error("fail to send sms: ", err.Error())
		return err
	}

	signInData := backend.ConfirmSmsVerifyCodeReq{
		Mobile: u.Data.Mobile,
		Code:   getVerifyCode(u.Data.Mobile),
	}
	token, err := backend.New().MobileSignIn(signInData)
	if err != nil {
		log.Error("fail to sign in: ", err.Error())
		return err
	}
	u.Token = token
	log.Debug("finish user sign in")
	return nil
}

func getVerifyCode(mobile string) (verifyCode string) {
	log.Debug("get verify code for mobile: ", mobile)
	verifyCode = t7Redis.New().Get(fmt.Sprintf("%s:%s", t7Redis.VerifyCodePrefix, mobile)).Val()
	log.Debug("got verify code: ", verifyCode)
	return
}
