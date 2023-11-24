package backend

import (
	"bytes"
	"encoding/json"
	"github.com/Template7/common/structs"
	"net/http"
)

func (c *client) MobileSignIn(data ConfirmSmsVerifyCodeReq) (token structs.Token, err error) {
	log.Debug("mobile sign in")

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriMobileSignIn, bytes.NewBuffer(bodyBytes))
	resp, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}

	log.Debug("sign in user response: ", string(resp))
	err = json.Unmarshal(resp, &token)
	return
}
