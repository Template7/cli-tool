package backend

import (
	"bytes"
	"encoding/json"
	"github.com/Template7/common/structs"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c client) CreateUser(data structs.User) (err error) {
	log.Debug("create user")

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriAdminCreateUser, bytes.NewBuffer(bodyBytes))
	resp, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}
	log.Debug("create user response: ", string(resp))
	return
}
