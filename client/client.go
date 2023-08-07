package client

import (
	"errors"

	"github.com/go-resty/resty/v2"

	"github.com/lawyzheng/go-fusion-compute/internal/common"
)

type Session string

type FusionComputeClient interface {
	Connect() error
	DisConnect() error
	SetSession(token string)
	GetSession() Session
	GetHost() string
	GetUser() string
	GetPassword() string
	GetHTTPClient() *resty.Client
	GetApiClient() (*resty.Client, error)
}

type RestyConstructor func() *resty.Client

func newDefaultConfig() *Config {
	return &Config{
		RestyConstructor: common.NewHttpClient,
	}
}

type Config struct {
	RestyConstructor RestyConstructor
}

func (c *Config) Merge(conf *Config) {
	if conf == nil {
		return
	}
	if conf.RestyConstructor != nil {
		c.RestyConstructor = conf.RestyConstructor
	}
}

func NewFusionComputeClient(host string, user string, password string, cfg ...*Config) FusionComputeClient {
	c := newDefaultConfig()
	for i := range cfg {
		c.Merge(cfg[i])
	}

	return &fusionComputeClient{
		user:     user,
		password: password,
		host:     host,
		config:   c,
	}
}

type fusionComputeClient struct {
	session  Session
	user     string
	password string
	host     string

	config *Config
}

func (f *fusionComputeClient) SetSession(token string) {
	f.session = Session(token)
}

func (f *fusionComputeClient) GetSession() Session {
	return f.session
}

func (f *fusionComputeClient) Connect() error {
	a := NewAuth(f)
	err := a.Login()
	if err != nil {
		return err
	}
	return nil
}

func (f *fusionComputeClient) DisConnect() error {
	a := NewAuth(f)
	err := a.Logout()
	if err != nil {
		return err
	}
	return nil
}
func (f *fusionComputeClient) GetHost() string {
	return f.host
}
func (f *fusionComputeClient) GetUser() string {
	return f.user
}
func (f *fusionComputeClient) GetPassword() string {
	return f.password
}

func (f *fusionComputeClient) GetHTTPClient() *resty.Client {
	return f.config.RestyConstructor()
}

func (f *fusionComputeClient) GetApiClient() (*resty.Client, error) {
	r := f.GetHTTPClient()
	if f.GetSession() == "" {
		return nil, errors.New("no session exists,please login and try it again")
	}
	f.setDefaultHeader(r)
	r.SetHeader(XAuthToken, string(f.GetSession())).
		SetHostURL(f.host)
	return r, nil
}

func (f *fusionComputeClient) setDefaultHeader(client *resty.Client) {
	client.SetHeaders(map[string]string{
		"Accept":          "application/json;version=v8.0;charset=UTF-8;",
		"Accept-Language": "zh_CN:1.0",
	})
}
