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

func newClient(cert, key, accountNumber string) (*http.Client, error) {
	t, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	originalTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{t},
		},
	}

	client := &http.Client{
		Transport: &customTransport{
			transport:     originalTransport,
			accountHeader: accountNumber,
		},
	}

	return client, nil
}

func sendRequest(client *http.Client, method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
