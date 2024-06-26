package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

func (c *Client) UpdateUserInfo(ctx context.Context, nickname string, userToken string) (err error) {
	log := c.log.WithContext(ctx).With("token", userToken)
	log.Debug("update user info")

	body := types.HttpUpdateUserInfoReq{
		Nickname: nickname,
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.endPoint, uriUpdateUserInfo), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return
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

	log.With("requestId", data.RequestId).Debug("update user info success")
	return nil
}

func (c *Client) GetUserInfo(ctx context.Context, userToken string) (types.HttpUserInfoRespData, error) {
	log := c.log.WithContext(ctx).With("token", userToken)
	log.Debug("get user info")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.endPoint, uriGetUserInfo), nil)
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return types.HttpUserInfoRespData{}, err
	}
	req.Header.Set("Authorization", userToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return types.HttpUserInfoRespData{}, err
	}

	var data types.HttpUserInfoResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return types.HttpUserInfoRespData{}, err
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return data.Data, nil
	}

	log.With("requestId", data.RequestId).Debug("get user info success")
	return data.Data, nil
}

func (c *Client) CreateUser(ctx context.Context, username string, password string, role string, nickname string, email string, adminToken string) error {
	log := c.log.WithContext(ctx).With("token", adminToken)
	log.Debug("create user")

	body := types.HttpCreateUserReq{
		Username: username,
		Password: password,
		Role:     role,
		Nickname: nickname,
		Email:    email,
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, uriCreateUser), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return err
	}
	req.Header.Set("Authorization", adminToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return err
	}

	var data types.HttpCreateUserResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).With("data", string(resp)).Error("fail to unmarshal data")
		return err
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	if !c.activateUser(ctx, data.Data.UserId, data.Data.ActivationCode, adminToken) {
		log.Warn("user activation fail")
		return fmt.Errorf("user activation fail")
	}

	log.Info("create user success and activated")
	return nil
}

func (c *Client) activateUser(ctx context.Context, userId string, actCode string, adminToken string) bool {
	log := c.log.WithContext(ctx).With("userId", userId)
	log.Debug("activate user")

	body := types.HttpActivateUserReq{
		ActivationCode: actCode,
	}
	bodyBytes, _ := json.Marshal(body)

	uri := fmt.Sprintf(uriActivateUser, userId)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, uri), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return false
	}
	req.Header.Set("Authorization", adminToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return false
	}

	var data types.HttpActivateUserResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).With("data", string(resp)).Error("fail to unmarshal data")
		return false
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return false
	}

	return data.Data.Success
}

func (c *Client) DeleteUser(ctx context.Context, userId string, adminToken string) error {
	log := c.log.WithContext(ctx).With("token", adminToken).With("userId", userId)
	log.Debug("delete user")

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.endPoint, fmt.Sprintf(uriDeleteUser, userId)), nil)
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return err
	}
	req.Header.Set("Authorization", adminToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return err
	}

	var data types.HttpRespBase
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).With("data", string(resp)).Error("fail to unmarshal data")
		return err
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.Info("delete user success")
	return nil
}
