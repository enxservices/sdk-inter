package oauth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	intersdk "github.com/enxservices/sdk-inter"
	"github.com/enxservices/sdk-inter/internal/types"
)

type OAuth struct {
	client *http.Client

	oauthData
}

type oauthData struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OauthResponse struct {
	AccessToken string           `json:"access_token"`
	TokenType   string           `json:"token_type"`
	ExpiresIn   int              `json:"expires_in"`
	Scope       []intersdk.Scope `json:"scope"`
}

func NewOAuth(client *http.Client, clientId, clientSecret string) *OAuth {
	return &OAuth{
		client: client,
		oauthData: oauthData{
			ClientID:     clientId,
			ClientSecret: clientSecret,
		},
	}
}

// Authorize authorizes the client with the provided scopes
func (o *OAuth) Authorize(scopes []intersdk.Scope) (*OauthResponse, error) {
	var resp OauthResponse

	var sn []string
	for _, scope := range scopes {
		sn = append(sn, scope.String())
	}

	form := url.Values{}
	form.Add("client_id", o.ClientID)
	form.Add("client_secret", o.ClientSecret)
	form.Add("grant_type", types.GrantType)
	form.Add("scope", strings.Join(sn, " "))

	req, err := http.NewRequest(http.MethodPost, types.OauthUrl, bytes.NewBufferString(form.Encode()))
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
		return nil, intersdk.ErrOauthFailed
	}

	// unmarshal
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetAccessToken returns the access token for the provided scopes (short function)
func (o *OAuth) GetAccessToken(scopes []intersdk.Scope) string {
	f, err := o.Authorize(scopes)
	if err != nil {
		return ""
	}

	return f.AccessToken
}
