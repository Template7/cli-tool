package user

import (
	backend2 "cli-tool/internal/backend"
)

func (u *User) SignUp() error {
	log.Debug("user sign up: ", u.Data.UserId)

	if err := backend2.New().SendSms(u.Data.Mobile); err != nil {
		log.Error("fail to send sms: ", err.Error())
		return err
	}

	signUpData := backend2.ConfirmSmsVerifyCodeReq{
		Mobile: u.Data.Mobile,
		Code:   getVerifyCode(u.Data.Mobile),
	}
	token, err := backend2.New().MobileSignUp(signUpData)
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
