package intersdk

type Inter struct {
	ClientID     string
	ClientSecret string
	KeyFilePath  string
	CertFilePath string

	// Client is the HTTP client
	Client *Client
	Oauth  *OAuth
}

type Option func(Inter)

// New creates a new Inter instance with the provided options
func New(options ...Option) Inter {
	i := Inter{}

	for _, option := range options {
		option(i)
	}

	// create a new HTTP client
	c, err := NewClient(i.CertFilePath, i.KeyFilePath)
	if err != nil {
		panic(err)
	}
	i.Client = c

	// create a new OAuth instance
	o := NewOAuth(i.Client, i.ClientID, i.ClientSecret)
	i.Oauth = o

	return i
}

// WithClientID sets the client ID for the Inter instance
func WithClientID(clientID string) Option {
	return func(i Inter) {
		i.ClientID = clientID
	}
}

// WithClientSecret sets the client secret for the Inter instance
func WithClientSecret(clientSecret string) Option {
	return func(i Inter) {
		i.ClientSecret = clientSecret
	}
}

// WithKeyFilePath sets the key file path for the Inter instance
func WithKeyFilePath(keyFilePath string) Option {
	return func(i Inter) {
		i.KeyFilePath = keyFilePath
	}
}

// WithCertFilePath sets the cert file path for the Inter instance
func WithCertFilePath(certFilePath string) Option {
	return func(i Inter) {
		i.CertFilePath = certFilePath
	}
}
