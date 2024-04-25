package intersdk

import (
	"errors"
	"net/http"

	"github.com/enxservices/sdk-inter/internal/oauth"
)

var (
	ErrTlsCertificateNil = errors.New("tls certificate not provided")
)

type Inter interface {
	// Charges - Boleto with Pix QR Code
	CreateCharge(charge Charge) (*Charge, error)
	GetCharge(uuid string) (*Charge, error)
	DowloadCharge(uuid string) ([]byte, error)
	CancelCharge(uuid string, reason string) error
}

type inter struct {
	ClientID      string
	ClientSecret  string
	ContaCorrente string

	client *http.Client
	Oauth  *oauth.OAuth
}

type Option func(*inter)

// New creates a new Inter instance with the provided key file path, certificate file path, client id and client secret
func New(keyFilePath, certFilePath, clientID, clientSecret, accountNumber string, options ...Option) (*inter, error) {
	i := &inter{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	for _, option := range options {
		option(i)
	}

	c, err := newClient(certFilePath, keyFilePath, accountNumber)
	if err != nil {
		return nil, err
	}

	i.client = c

	o := oauth.NewOAuth(c, clientID, clientSecret)
	i.Oauth = o

	return i, nil
}
