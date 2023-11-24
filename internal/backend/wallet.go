package backend

import (
	"bytes"
	"encoding/json"
	"github.com/Template7/backend/pkg/apiBody"
	"net/http"
)

func (c *client) Deposit(data apiBody.DepositReq, userToken string) (err error) {
	log.Debug("backend deposit")

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriDeposit, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", userToken)
	res, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}

	log.Debug("deposit response: ", string(res))
	return
}

func (c *client) Transaction(data apiBody.TransactionReq, userToken string) (resp apiBody.TransactionResp, err error) {
	log.Debug("make transaction: ", data.String())

	bodyBytes, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, c.endPoint+uriTransaction, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", userToken)
	res, httpErr := c.SendReq(req)
	if httpErr != nil {
		err = httpErr
		log.Error("fail to send http request: ", httpErr.Error())
		return
	}

	log.Debug("make transaction response: ", string(res))
	err = json.Unmarshal(res, &resp)
	return
}
