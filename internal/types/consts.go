package types

const (
	BaseUrlProduction      = "https://cdpj.partners.bancointer.com.br"
	BaseUrlSandBox         = "https://cdpj-sandbox.partners.uatinter.co"
	OauthUrl               = "oauth/v2/token"
	CobPixBoletoUrl        = "cobranca/v3/cobrancas"
	CobBoletoUrl           = "cobranca/v2/boletos"
	CobPixUrl              = "pix/v2"
	CobWebHookUrl          = "cobranca/v3/cobrancas/webhook"
	CobWebHookUrlCallbacks = "cobranca/v3/cobrancas/webhook/callbacks"
	GrantType              = "client_credentials"
)

type Env string

const (
	EnvProd    Env = "production"
	EnvSandbox Env = "sandbox"
)
