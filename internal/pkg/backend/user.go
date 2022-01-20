package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/pkg/apiBody"
	"github.com/Template7/common/structs"
	"net/http"
)

func (c client) CreateUser(data apiBody.CreateUserReq) (err error) {
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

func (c client) UpdateUser(data structs.User, userToken string) (err error) {
	log.Debug("update user")

	bodyBytes, _ := json.Marshal(data.BasicInfo)

	uri := fmt.Sprintf(uriUpdateUser, data.UserId)
	req, _ := http.NewRequest(http.MethodPut, c.endPoint+uri, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", userToken)
	resp, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}
	log.Debug("create user response: ", string(resp))
	return
}
