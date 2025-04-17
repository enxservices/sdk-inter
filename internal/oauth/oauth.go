package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	types2 "github.com/enxservices/sdk-inter/internal/types"
)

type OAuth struct {
	client     *http.Client
	tokenStore map[types2.Scope]*OauthResponse
	baseURL    string
	oauthData
}

type oauthData struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OauthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   time.Time
}

func NewOAuth(client *http.Client, clientId, clientSecret string, baseURL string) *OAuth {
	return &OAuth{
		client:     client,
		tokenStore: make(map[types2.Scope]*OauthResponse),
		baseURL:    baseURL,
		oauthData: oauthData{
			ClientID:     clientId,
			ClientSecret: clientSecret,
		},
	}
}

// Authorize authorizes the client with the provided scopes
func (o *OAuth) Authorize(scope types2.Scope) (*OauthResponse, error) {
	var resp OauthResponse

	form := url.Values{}
	form.Add("client_id", o.ClientID)
	form.Add("client_secret", o.ClientSecret)
	form.Add("grant_type", types2.GrantType)
	form.Add("scope", scope.String())

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", o.baseURL, types2.OauthUrl), bytes.NewBufferString(form.Encode()))
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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha ao realizar oauth: %s", body)
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	resp.CreatedAt = time.Now()

	return &resp, nil
}

// margem de erro de 20 min, portanto ao inves de validar token com datacriacao <= 60 min validamos com 40 min
func (o *OAuth) isValidToken(token *OauthResponse) bool {
	expirationTime := token.CreatedAt.Add(time.Duration(token.ExpiresIn) * time.Second)
	safetyMargin := 20 * time.Minute
	return time.Now().Before(expirationTime.Add(-safetyMargin))
}

// GetAccessToken returns the access token for the provided scopes (short function)
func (o *OAuth) GetAccessToken(scope types2.Scope) (string, error) {
	if token, exists := o.tokenStore[scope]; exists {
		if o.isValidToken(token) {
			return token.AccessToken, nil
		}
	}

	token, err := o.Authorize(scope)
	if err != nil {
		return "", err
	}

	o.tokenStore[scope] = token

	return token.AccessToken, nil
}
