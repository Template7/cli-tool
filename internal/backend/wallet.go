package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"github.com/Template7/backend/pkg/apiBody"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c *CliCent) Deposit(ctx context.Context, walletId string, currency string, amount int, userToken string) error {
	log := c.log.WithContext(ctx).With("currency", currency).With("amount", amount).With("token", userToken)
	log.Debug("wallet deposit")

	body := types.HttpWalletDepositReq{
		Currency: currency,
		Amount:   uint32(amount),
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, fmt.Sprintf(uriWalletDeposit, walletId)), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return err
	}
	req.Header.Set("Authorization", userToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return err
	}

	var data types.HttpRespBase
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return err
	}

	log.With("requestId", data.RequestId).Debug("wallet deposit success")
	return nil
}

func (c *CliCent) Withdraw(ctx context.Context, walletId string, currency string, amount int, userToken string) error {
	log := c.log.WithContext(ctx).With("currency", currency).With("amount", amount).With("token", userToken)
	log.Debug("wallet withdraw")

	body := types.HttpWalletWithdrawReq{
		Currency: currency,
		Amount:   uint32(amount),
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, fmt.Sprintf(uriWalletWithdraw, walletId)), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return err
	}
	req.Header.Set("Authorization", userToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return err
	}

	var data types.HttpRespBase
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return err
	}

	log.With("requestId", data.RequestId).Debug("wallet withdraw success")
	return nil
}

func (c *CliCent) Transaction(data apiBody.TransactionReq, userToken string) (resp apiBody.TransactionResp, err error) {
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
