package intersdk

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	*http.Client
}

// NewClient creates a new client with the provided tls certificate.
func NewClient(cert, key string) (*Client, error) {
	t, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{t},
		},
	}

	return &Client{
		Client: &http.Client{Transport: tr},
	}, nil
}
