package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

func (c *Client) NativeLogin(ctx context.Context, username string, password string) string {
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

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.Warn("login fail")
		return ""
	}

	log.Info("user login success")
	return data.Data.Token
}
