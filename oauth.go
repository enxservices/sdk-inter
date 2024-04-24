package intersdk

type OAuth struct {
	client *Client
}

type OauthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
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

func (o *OAuth) Authorize(req *OauthRequest) (*OauthResponse, error) {
	var resp OauthResponse

	return &resp, nil
}
