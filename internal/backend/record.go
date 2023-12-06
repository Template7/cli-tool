package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Template7/backend/api/types"
	"net/http"
)

func (c *Client) GetWalletBalanceRecord(ctx context.Context, walletId string, currency string, userToken string) []types.HttpGetWalletBalanceRecordRespData {
	log := c.log.WithContext(ctx).With("walletId", walletId).With("currency", currency)
	log.Debug("get wallet balance record")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.endPoint, fmt.Sprintf(uriGetWalletBalanceRecord, walletId, currency)), nil)
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

	var data types.HttpGetWalletBalanceRecordResp
	if err := json.Unmarshal(resp, &data); err != nil {
		log.WithError(err).Error("fail to unmarshal data")
		return nil
	}

	log = log.With("requestId", data.RequestId)

	if data.Code != types.HttpRespCodeOk {
		log.With("resp", resp).Warn("something went wrong")
		return nil
	}

	log.Debug("get wallet balance record success")
	return data.Data
}
