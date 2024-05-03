package intersdk

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

type customTransport struct {
	transport     http.RoundTripper
	accountHeader string
}

func (c *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-conta-corrente", c.accountHeader)
	return c.transport.RoundTrip(req)
}

func NewClient(cert, key string, accountNumber *string) (*http.Client, error) {
	t, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	transport := createTransport(t)

	client := &http.Client{
		Transport: createCustomTransport(transport, accountNumber),
	}

	return client, nil
}

func createTransport(cert tls.Certificate) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
}

func createCustomTransport(transport *http.Transport, accountNumber *string) http.RoundTripper {
	if accountNumber == nil {
		return &customTransport{
			transport: transport,
		}
	}
	return &customTransport{
		transport:     transport,
		accountHeader: *accountNumber,
	}
}

func sendRequest(client *http.Client, method, url, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
