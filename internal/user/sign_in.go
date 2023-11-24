package user

import (
	backend2 "cli-tool/internal/backend"
)

func (u *User) SignIn() error {
	log.Debug("user sign in: ", u.Data.UserId)

	if err := backend2.New().SendSms(u.Data.Mobile); err != nil {
		log.Error("fail to send sms: ", err.Error())
		return err
	}

	signInData := backend2.ConfirmSmsVerifyCodeReq{
		Mobile: u.Data.Mobile,
		Code:   getVerifyCode(u.Data.Mobile),
	}
	token, err := backend2.New().MobileSignIn(signInData)
	if err != nil {
		log.Error("fail to sign in: ", err.Error())
		return err
	}
	if err := u.updateToken(token); err != nil {
		log.Error("fail to update token: ", err.Error())
		return err
	}

	log.Debug("finish user sign in")
	return nil
}

// GetInfo
// get user info and save into the property
func (u *User) GetInfo() error {
	log.Debug("get user info: ", u.Data.UserId)

	data, err := backend2.New().GetUserInfo(u.Data.UserId, u.Token.AccessToken)
	if err != nil {
		log.Error("fail to get user data: ", u.Data.UserId, ". ", err.Error())
		return err
	}

	u.Data.BasicInfo = data.UserInfo
	u.Wallet = data.WalletData

	log.Debug("finish get user info: ", u.Data.UserId)
	return nil
}
