package util

import (
	t7Error2 "cli-tool/internal/t7Error"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func SendHttpRequest(req *http.Request) (response []byte, code int, err *t7Error2.Error) {
	client := http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		log.Error("fail to send request: ", httpErr.Error())
		err = t7Error2.HttpOperationFail.WithDetailAndStatus(httpErr.Error(), http.StatusInternalServerError)
		return
	}

	response, _ = ioutil.ReadAll(resp.Body)
	code = resp.StatusCode

	//if resp.StatusCode != http.StatusOK {
	//	log.Info("unexpected response: ", string(response))
	//	log.Error("unexpected response code: ", resp.StatusCode)
	//	err = t7Error.HttpUnexpectedResponseCode.WithDetailAndStatus(resp.Status, http.StatusInternalServerError)
	//	return
	//}
	return
}
