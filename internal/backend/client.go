package backend

import (
	"cli-tool/internal/config"
	"context"
	"github.com/Template7/common/logger"
	"io"
	"net/http"
	"sync"
)

const (
	uriUserLogin      = "/api/v1/login/native"
	uriUpdateUserInfo = "/api/v1/user/info"
	uriGetUserInfo    = "/api/v1/user/info"

	uriWalletDeposit  = "/api/v1/wallets/%s/deposit"
	uriWalletWithdraw = "/api/v1/wallets/%s/withdraw"
	uriWalletTransfer = "/api/v1/transfer"
	//uriSendSmsVerifyCode = "/api/v1/verify-code/sms"
	//uriMobileSignUp      = "/api/v1/sign-up/mobile"
	//uriMobileSignIn      = "/api/v1/sign-in/mobile"
	//uriUpdateUser        = "/api/v1/users/%s"
	//uriGetUserData       = "/api/v1/users/%s"
	//uriDeposit           = "/api/v1/wallet/deposit"
	//uriTransaction       = "/api/v1/transaction"
)

type CliCent struct {
	endPoint string

	log *logger.Logger
}

var (
	once     sync.Once
	instance *CliCent
)

func New() *CliCent {
	once.Do(func() {
		log := logger.New().WithService("backend")
		instance = &CliCent{
			endPoint: config.New().Backend.Endpoint,
			log:      log,
		}

		log.Debug("backend client initialized")
	})
	return instance
}

func (c *CliCent) SendReq(ctx context.Context, req *http.Request) (data []byte, err error) {
	log := c.log.WithContext(ctx)
	log.Debug("send http req")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("fail to read resp body")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.With("status", resp.StatusCode).With("resp", string(data)).Warn("non-200 http status")
	}

	return data, nil
}
