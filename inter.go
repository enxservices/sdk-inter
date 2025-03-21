package intersdk

import (
	"errors"
	"net/http"

	"github.com/enxservices/sdk-inter/internal/oauth"
	"github.com/enxservices/sdk-inter/internal/types"
)

var ErrTlsCertificateNil = errors.New("tls certificate not provided")

type Inter interface {
	CreateCharge(charge CreateChargeRequest) (string, error)
	GetCharge(solicitationCode string) (*ChargeResponse, error)
	GetChargeList(params QueryParamChargeList) (*ChargeList, error)
	DowloadCharge(solicitationCode string) (string, error)
	CancelCharge(solicitationCode string, reason string) error
	GetWebhook() (*Webhook, error)
	DeleteWebhook() (*WebhookError, error)
	CreateWebhook(webhookUrl string) error
}

type inter struct {
	ClientID      string
	ClientSecret  string
	ContaCorrente string
	BaseURL       string
	DisableTLS    bool

	client *http.Client
	Oauth  *oauth.OAuth
}

type Option func(*inter)

func WithSandboxEnv() Option {
	return func(i *inter) {
		i.BaseURL = types.BaseUrlSandBox
	}
}

func WithLocalEnv(url string) Option {
	return func(i *inter) {
		i.BaseURL = url
	}
}

func WithDisableTLS() Option {
	return func(i *inter) {
		i.DisableTLS = true
	}
}

// New creates a new Inter instance with the provided key file path, certificate file path, client id and client secret
func New(keyFilePath, certFilePath, clientID, clientSecret string, accountNumber *string, options ...Option) (Inter, error) {
	i := &inter{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      types.BaseUrlProduction,
	}

	for _, option := range options {
		option(i)
	}

	var err error
	if !i.DisableTLS {
		i.client, err = NewClient(certFilePath, keyFilePath, accountNumber)
		if err != nil {
			return nil, err
		}
	} else {
		i.client = http.DefaultClient
	}

	o := oauth.NewOAuth(i.client, clientID, clientSecret, i.BaseURL)
	i.Oauth = o

	return i, nil
}
