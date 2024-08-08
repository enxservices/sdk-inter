package intersdk

import (
	"errors"
	"net/http"

	"github.com/enxservices/sdk-inter/internal/oauth"
	"github.com/enxservices/sdk-inter/internal/types"
)

var (
	ErrTlsCertificateNil = errors.New("tls certificate not provided")
)

type Inter interface {
	// Charges - Boleto with Pix QR Code
	CreateCharge(charge CreateChargeRequest) (string, error)
	GetCharge(solicitationCode string) (*ChargeResponse, error)
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
	Environment   string

	client *http.Client
	Oauth  *oauth.OAuth
}

type Option func(*inter)

func WithEnvironment(environment string) Option {
	return func(i *inter) {
		switch environment {
		case "sandbox":
			i.Environment = types.BaseUrlSandBox
		case "production":
			i.Environment = types.BaseUrlProduction
		default:
			panic("invalid environment provided, please use 'sandbox' or 'production'")
		}
	}
}

// New creates a new Inter instance with the provided key file path, certificate file path, client id and client secret
func New(environment, keyFilePath, certFilePath, clientID, clientSecret string, accountNumber *string, options ...Option) (Inter, error) {
	i := &inter{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Environment:  types.BaseUrlProduction,
	}

	for _, option := range options {
		option(i)
	}

	c, err := NewClient(certFilePath, keyFilePath, accountNumber)
	if err != nil {
		return nil, err
	}

	i.client = c

	o := oauth.NewOAuth(c, clientID, clientSecret)
	i.Oauth = o

	return i, nil
}
