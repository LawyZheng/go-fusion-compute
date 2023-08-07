package common

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	fcErr "github.com/lawyzheng/go-fusion-compute/pkg/error"
)

func NewHttpClient() *resty.Client {
	r := resty.New()
	r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return r
}

func FormatHttpError(resp *resty.Response, err fcErr.Error) error {
	err.SetHTTPStatus(resp.StatusCode())
	if e := json.Unmarshal(resp.Body(), err); e != nil {
		return fmt.Errorf("status=%d;body=%s", resp.StatusCode(), string(resp.Body()))
	}
	return err
}
