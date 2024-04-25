package intersdk

type Scope string

// Scopes are used to define the permissions that the application has to access the API.
const (
	ExtratoRead         Scope = "extrato.read"
	BoletoCobrancaRead  Scope = "boleto-cobranca.read"
	BoletoCobrancaWrite Scope = "boleto-cobranca.write"
	CobWrite            Scope = "cob.write"
	CobRead             Scope = "cob.read"
	CobvRead            Scope = "cobv.read"
	CobvWrite           Scope = "cobv.write"
	PixRead             Scope = "pix.read"
	PixWrite            Scope = "pix.write"
	WebhookRead         Scope = "webhook.read"
	WebhookWrite        Scope = "webhook.write"
)

func (s Scope) String() string {
	return string(s)
}
