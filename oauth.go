package intersdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type OAuth struct {
	client *Client
}

type OauthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scopes       string `json:"scope"`
}

type OauthResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiresIn   int     `json:"expires_in"`
	Scope       []Scope `json:"scope"`
}

func NewOAuth(client *Client) *OAuth {
	return &OAuth{
		client: client,
	}
}

func (o *OAuth) Authorize(data *OauthRequest) (*OauthResponse, error) {
	var resp OauthResponse

	form := url.Values{}
	form.Add("client_id", data.ClientID)
	form.Add("client_secret", data.ClientSecret)
	form.Add("grant_type", data.GrantType)
	form.Add("scope", data.Scopes)

	req, err := http.NewRequest(http.MethodPost, OauthUrl, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if body == nil || res.StatusCode != http.StatusOK {
		return nil, ErrOauthFailed
	}

	// unmarshal
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
