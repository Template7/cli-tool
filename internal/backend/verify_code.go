package backend

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// TODO: ref to backend
type SmsReq struct {
	Mobile string `json:"mobile" binding:"required" example:"+886987654321"`
}

func (c *client) SendSms(mobile string) (err error) {
	log.Debug("send sms verify code: ", mobile)

	data := SmsReq{
		Mobile: mobile,
	}
	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriSendSmsVerifyCode, bytes.NewBuffer(bodyBytes))
	resp, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}
	log.Debug("send sms response: ", string(resp))
	return
}
