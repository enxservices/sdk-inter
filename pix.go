package intersdk

type Calendar struct {
	Expiration int      `json:"expiracao"`
	Customer   Customer `json:"devedor"`
}

// Charge Type
type LocType string

const (
	CobV LocType = "cobv"
	COB  LocType = "cob"
)

type Loc struct {
}

type Customer interface {
	isCustomer()
}

type PessoaFisica struct {
	CPF  string `json:"cpf"`
	Nome string `json:"nome"`
}

func (pf PessoaFisica) isCustomer() {}

type PessoaJuridica struct {
	CNPJ string `json:"cnpj"`
	Nome string `json:"nome"`
}

func (pj PessoaJuridica) isCustomer() {}

type CreateImmediateCharge struct {
	Calendar Calendar `json:"calendario"`
	Loc      Loc
}
