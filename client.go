package intersdk

import (
	"crypto/tls"
	"net/http"
)

// NewClient creates a new client with the provided tls certificate.
func newClient(cert, key string) (*http.Client, error) {
	t, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{t},
		},
	}

	return &http.Client{Transport: tr}, nil
}
