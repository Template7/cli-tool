package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

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

func (c *CliCent) GetUserInfo(ctx context.Context, userToken string) (types.HttpUserInfoRespData, error) {
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

	log.With("requestId", data.RequestId).Debug("get user info success")
	return data.Data, nil
}
