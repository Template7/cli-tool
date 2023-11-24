package backend

import (
	"bytes"
	"encoding/json"
	"github.com/Template7/common/structs"
	"net/http"
)

// TODO: ref to backend
type ConfirmSmsVerifyCodeReq struct {
	Mobile string `json:"mobile" binding:"required" example:"+886987654321"`
	Code   string `json:"code" binding:"required" example:"1234567"`
}

func (c *client) MobileSignUp(data ConfirmSmsVerifyCodeReq) (token structs.Token, err error) {
	log.Debug("mobile sign up")

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriMobileSignUp, bytes.NewBuffer(bodyBytes))
	resp, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}

	log.Debug("sign up user response: ", string(resp))
	err = json.Unmarshal(resp, &token)
	return
}
