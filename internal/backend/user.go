package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

func (c *CliCent) NativeLogin(ctx context.Context, username string, password string) string {
	log := c.log.WithContext(ctx).With("username", username)
	log.Debug("user login")

	body := types.HttpLoginReq{
		Username: username,
		Password: password,
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.endPoint, uriUserLogin), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return ""
	}
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return ""
	}

	var data types.HttpLoginResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return ""
	}

	log.With("requestId", data.RequestId).Debug("user login success")
	return data.Data.Token
}

func (c *CliCent) UpdateUserInfo(ctx context.Context, nickname string, userToken string) (err error) {
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

	log.With("requestId", data.RequestId).Debug("update user info success")
	return nil
}

func (c *CliCent) GetUserInfo(ctx context.Context, userToken string) (types.HttpUserInfoResp, error) {
	log := c.log.WithContext(ctx).With("token", userToken)
	log.Debug("get user info")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.endPoint, uriGetUserInfo), nil)
	if err != nil {
		log.WithError(err).Error("fail to new http req")
		return types.HttpUserInfoResp{}, err
	}
	req.Header.Set("Authorization", userToken)
	resp, err := c.SendReq(ctx, req)
	if err != nil {
		log.WithError(err).Error("fail to send req")
		return types.HttpUserInfoResp{}, err
	}

	var data types.HttpUserInfoResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return types.HttpUserInfoResp{}, err
	}

	log.With("requestId", data.RequestId).Debug("get user info success")
	return data, nil
}
