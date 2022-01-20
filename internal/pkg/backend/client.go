package backend

import (
	"bytes"
	"cli-tool/internal/pkg/config"
	"cli-tool/internal/pkg/t7Error"
	"cli-tool/internal/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/Template7/common/logger"
	"github.com/Template7/common/structs"
	"net/http"
	"sync"
)

const (
	uriAdminSignIn     = "/admin/v1/sign-in"
	uriAdminCreateUser = "/admin/v1/user"

	uriSendSmsVerifyCode = "/api/v1/verify-code/sms"
	uriMobileSignUp      = "/api/v1/sign-up/mobile"
	uriMobileSignIn      = "/api/v1/sign-in/mobile"
	uriUpdateUser        = "/api/v1/users/%s"
)

var (
	log = logger.GetLogger()
)

type client struct {
	endPoint string
	username string
	password string
	token    string
}

var (
	once     sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		instance = &client{
			endPoint: config.New().Backend.Endpoint,
			username: config.New().Backend.Username,
			password: config.New().Backend.Password,
		}
		//instance.SignIn()
		log.Debug("backend client initialized")
	})
	return instance
}

func (c *client) SignIn() {
	log.Debug("sign in admin")

	body := structs.Admin{
		Username: c.username,
		Password: c.password,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriAdminSignIn, bytes.NewBuffer(bodyBytes))
	resp, err := c.SendReq(req)
	if err != nil {
		log.Fatal("admin sign fail: ", err.Error())
	}

	var token structs.Token
	if err := json.Unmarshal(resp, &token); err != nil {
		log.Fatal("admin sign fail: ", err.Error())
	}
	c.token = token.AccessToken
	log.Debug("admin sign in successfully")
	return
}

func (c client) SendReq(req *http.Request) (response []byte, err *t7Error.Error) {
	//req.Header.Set("Authorization", c.token)
	resp, code, httpErr := util.SendHttpRequest(req)
	if httpErr != nil {
		return nil, httpErr
	}
	if code >= 400 {
		err = t7Error.HttpUnexpectedResponseCode.WithDetail(fmt.Sprintf("status code: %d", code))
	}

	return resp, nil
}
