package backend

import (
	"bytes"
	"cli-tool/internal/pkg/db/collection"
	"cli-tool/internal/pkg/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c client) CreateUser(data collection.User) (err error) {
	log.Debug("create user")

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriAdminSignIn, bytes.NewBuffer(bodyBytes))
	resp, httpErr := util.SendHttpRequest(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}
	log.Debug("create user response: ", string(resp))
	return
}
