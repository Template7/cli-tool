package backend

import (
	"bytes"
	"cli-tool/internal/pkg/config"
	"cli-tool/internal/pkg/db/collection"
	"cli-tool/internal/pkg/t7Error"
	"cli-tool/internal/pkg/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

const (
	uriAdminSignIn     = "/admin/v1/sign-in"
	uriAdminCreateUser = "/admin/v1/user"
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
		instance.SignIn()
		log.Debug("backend client initialized")
	})
	return instance
}

func (c *client) SignIn() {
	log.Debug("sign in admin")

	body := collection.Admin{
		Username: c.username,
		Password: c.password,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriAdminSignIn, bytes.NewBuffer(bodyBytes))
	resp, err := util.SendHttpRequest(req)
	if err != nil {
		log.Fatal("admin sign fail: ", err.Error())
	}

	var token collection.Token
	if err := json.Unmarshal(resp, &token); err != nil {
		log.Fatal("admin sign fail: ", err.Error())
	}
	c.token = token.AccessToken
	log.Debug("admin sign in successfully")
	return
}

func (c client) SendReq(req *http.Request) (response []byte, err *t7Error.Error) {
	req.Header.Set("Authorization", c.token)
	return util.SendHttpRequest(req)
}
