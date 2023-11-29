package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

func (c *Client) GetUserWallets(ctx context.Context, userToken string) []types.HttpGetUserWalletsRespData {
	log := c.log.WithContext(ctx).With("token", userToken)
	log.Debug("get user wallet")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.endPoint, uriGetUserWallets), nil)
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return nil
	}
	req.Header.Set("Authorization", userToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return nil
	}

	var data types.HttpGetUserWalletsResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return nil
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.With("requestId", data.RequestId).Debug("get user wallet success")
	return data.Data
}

func (c *Client) Deposit(ctx context.Context, walletId string, currency string, amount int, userToken string) error {
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

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.With("requestId", data.RequestId).Debug("wallet deposit success")
	return nil
}

func (c *Client) Withdraw(ctx context.Context, walletId string, currency string, amount int, userToken string) error {
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

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.With("requestId", data.RequestId).Debug("wallet withdraw success")
	return nil
}

func (c *Client) Transfer(ctx context.Context, fromWalletId string, toWalletId string, currency string, amount int, userToken string) error {
	log := c.log.WithContext(ctx).With("currency", currency).With("amount", amount).With("token", userToken).With("from", fromWalletId).With("to", toWalletId)
	log.Debug("wallet transfer")

	body := types.HttpTransferMoneyReq{
		FromWalletId: fromWalletId,
		ToWalletId:   toWalletId,
		Currency:     currency,
		Amount:       uint32(amount),
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, uriWalletTransfer), bytes.NewBuffer(bodyBytes))
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

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.Debug("wallet transfer success")
	return nil
}
