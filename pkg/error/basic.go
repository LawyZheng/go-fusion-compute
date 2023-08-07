package error

import "fmt"

type Error interface {
	error
	SetHTTPStatus(code int)
}

type Basic struct {
	HTTPStatus       int    `json:"-"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDes"`
}

func (b *Basic) SetHTTPStatus(code int) {
	b.HTTPStatus = code
}

func (b *Basic) Error() string {
	return fmt.Sprintf("status=%d;errcode=%s;msg=%s", b.HTTPStatus, b.ErrorCode, b.ErrorDescription)
}
