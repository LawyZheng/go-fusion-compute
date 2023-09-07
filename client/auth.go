package client

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/lawyzheng/go-fusion-compute/internal/common"
	fcErr "github.com/lawyzheng/go-fusion-compute/pkg/error"
)

const (
	XAuthUser     = "X-Auth-User"
	XAuthKey      = "X-Auth-Key"
	XAuthUserType = "X-Auth-UserType"
	XAuthToken    = "X-Auth-Token"

	authUri = "/service/session"
)

type Auth interface {
	Login(ctx context.Context) error
	Logout(ctx context.Context) error
}

func NewAuth(client FusionComputeClient) Auth {
	return &auth{client: client}
}

type auth struct {
	client FusionComputeClient
}

func (a *auth) Login(ctx context.Context) error {
	host := a.client.GetHost()
	r := a.client.GetHTTPClient()
	r.SetBaseURL(host).
		SetHeader(XAuthUser, a.client.GetUser()).
		SetHeader(XAuthKey, encodePassword(a.client.GetPassword())).
		SetHeader(XAuthUserType, "2")
	req := r.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	resp, err := req.Post(authUri)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		var loginResponse LoginResponse
		_ = json.Unmarshal(resp.Body(), &loginResponse)
		token := resp.Header().Get(XAuthToken)
		a.client.SetSession(token)
	} else {
		e := new(fcErr.Basic)
		return common.FormatHttpError(resp, e)
	}
	return nil
}

func (a *auth) Logout(ctx context.Context) error {
	host := a.client.GetHost()
	r := a.client.GetHTTPClient()
	r.SetBaseURL(host).
		SetHeader(XAuthToken, string(a.client.GetSession()))
	req := r.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	resp, err := req.Delete(authUri)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		a.client.SetSession("")
	} else {
		e := new(fcErr.Basic)
		return common.FormatHttpError(resp, e)
	}
	return nil
}

func encodePassword(pass string) string {
	bs := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(bs[:])
}
