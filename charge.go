package intersdk

import (
	"time"
)

type State string

const (
	StateAcre             State = "AC"
	StateAlagoas          State = "AL"
	StateAmapa            State = "AP"
	StateAmazonas         State = "AM"
	StateBahia            State = "BA"
	StateCeara            State = "CE"
	StateDistritoFederal  State = "DF"
	StateEspiritoSanto    State = "ES"
	StateGoias            State = "GO"
	StateMaranhao         State = "MA"
	StateMatoGrosso       State = "MT"
	StateMatoGrossoDoSul  State = "MS"
	StateMinasGerais      State = "MG"
	StatePara             State = "PA"
	StateParaiba          State = "PB"
	StateParana           State = "PR"
	StatePernambuco       State = "PE"
	StatePiaui            State = "PI"
	StateRioDeJaneiro     State = "RJ"
	StateRioGrandeDoNorte State = "RN"
	StateRioGrandeDoSul   State = "RS"
	StateRondonia         State = "RO"
	StateRoraima          State = "RR"
	StateSantaCatarina    State = "SC"
	StateSaoPaulo         State = "SP"
	StateSergipe          State = "SE"
	StateTocantins        State = "TO"
)

type CodeTypeDiscount string

const (
	ValueFixedInfoData CodeTypeDiscount = "VALORFIXODATAINFORMADA"
	PercentageInfoData CodeTypeDiscount = "PERCENTUALDATAINFORMADA"
)

type CodeTypeFee string

const (
	FeePercentage CodeTypeFee = "PERCENTUAL"
	FeeFix        CodeTypeFee = "VALORFIXO"
)

type CodeTypeLatePayment string

const (
	LatePaymentFeeDaily   CodeTypeLatePayment = "VALORDIA"
	LatePaymentFeeMonthly CodeTypeLatePayment = "TAXAMENSAL"
)

type KindPeople string

const (
	PJ KindPeople = "JURIDICA"
	PF KindPeople = "FISICA"
)

type Payer struct {
	Email      string     `json:"email"`
	DDD        string     `json:"ddd"`
	Phone      string     `json:"telefone"`
	Number     string     `json:"numero"`
	Complement string     `json:"complemento"`
	Doc        string     `json:"cpfCnpj"`
	KindPeople KindPeople `json:"tipoPessoa"`
	Name       string     `json:"nome"`
	Address    string     `json:"endereco"`
	District   string     `json:"bairro"`
	City       string     `json:"cidade"`
	State      State      `json:"uf"`
	ZipCode    string     `json:"cep"`
}

type Message struct {
	LineOne   *string `json:"linha1"`
	LineTwo   *string `json:"linha2"`
	LineThree *string `json:"linha3"`
	LineFour  *string `json:"linha4"`
	LineFive  *string `json:"linha5"`
}

type Discount struct {
	Fee          int              `json:"taxa"`
	Code         CodeTypeDiscount `json:"codigo"`
	DaysQuantity int              `json:"quantidadeDias"`
}

type Fine struct {
	Fee  int         `json:"taxa"`
	Code CodeTypeFee `json:"codigo"`
}

type LatePaymentFee struct {
	Fee  int                 `json:"taxa"`
	Code CodeTypeLatePayment `json:"codigo"`
}

type Beneficiary struct {
	Doc        string     `json:"cpfCnpj"`
	KindPeople KindPeople `json:"tipoPessoa"`
	Name       string     `json:"nome"`
	Address    string     `json:"endereco"`
	District   string     `json:"bairro"`
	City       string     `json:"cidade"`
	State      State      `json:"uf"`
	ZipCode    string     `json:"cep"`
}

// Date Format YYYY-MM-DD (Request STRUCT)
type CreateCharge struct {
	YourNumber       string          `json:"seuNumero"`
	NominalValue     float64         `json:"valorNominal"`
	DueDate          time.Time       `json:"dataVencimento"`
	DaysAfterDue     int             `json:"numDiasAgenda"`
	Payer            Payer           `json:"pagador"`
	Discount         *Discount       `json:"desconto,omitempty"`
	Fine             *Fine           `json:"multa,omitempty"`
	LatePayment      *LatePaymentFee `json:"juros,omitempty"`
	Message          *Message        `json:"mensagem,omitempty"`
	FinalBeneficiary Beneficiary     `json:"beneficiarioFinal"`
}

type Charge struct {
}

type Ticket struct {
}
