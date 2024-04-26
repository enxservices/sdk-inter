package intersdk

import (
	"errors"
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

type ChargeStatus string

const (
	ChargeStatusPending        ChargeStatus = "A_RECEBER"
	ChargeStatusPaid           ChargeStatus = "RECEBIDO"
	ChargeStatusCanceled       ChargeStatus = "CANCELADO"
	ChargeStatusExpired        ChargeStatus = "EXPIRADO"
	ChargeStatusLate           ChargeStatus = "ATRASADO"
	ChargeStatusMarkedReceived ChargeStatus = "MARCADO_RECEBIDO"
)

type ChargeType string

const (
	ChargeTypeSimple    ChargeType = "SIMPLES"
	ChargeTypeRecurrent ChargeType = "RECORRENTE"
	ChargeTypeSplit     ChargeType = "PARCELADO"
)

type SourceReceived string

const (
	SourceReceivedPIX    SourceReceived = "PIX"
	SourceReceivedBoleto SourceReceived = "BOLETO"
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
	Fee   int                 `json:"taxa"`
	Code  CodeTypeLatePayment `json:"codigo"`
	Valor *int                `json:"valor"`
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
type CreateChargeRequest struct {
	YourNumber       string          `json:"seuNumero"`
	NominalValue     float64         `json:"valorNominal"`
	DueDate          time.Time       `json:"dataVencimento"`
	DaysAfterDue     int             `json:"numDiasAgenda"`
	Payer            Payer           `json:"pagador"`
	Discount         *Discount       `json:"desconto,omitempty"`
	Fine             *Fine           `json:"multa,omitempty"`
	LatePayment      *LatePaymentFee `json:"mora,omitempty"`
	Message          *Message        `json:"mensagem,omitempty"`
	FinalBeneficiary Beneficiary     `json:"beneficiarioFinal"`
}

type Charge struct {
	SolicitationCode    string         `json:"codigoSolicitacao"`
	YourNumber          string         `json:"seuNumero"`
	EmissionDate        string         `json:"dataEmissao"`
	NominalValue        int            `json:"valorNominal"`
	DueDate             string         `json:"dataVencimento"`
	DaysAfterDue        int            `json:"numDiasAgenda"`
	ChargeType          ChargeType     `json:"tipoCobranca"`
	Status              ChargeStatus   `json:"situacao"`
	StatusDate          string         `json:"dataSituacao"`
	TotalAmountReceived int            `json:"valorTotalRecebido"`
	SourceReceived      SourceReceived `json:"origemRecebimento"`
	CancellationReason  string         `json:"motivoCancelamento"`
	Archived            bool           `json:"arquivada"`
	Discounts           []Discount     `json:"descontos"`
	Fine                Fine           `json:"multa"`
	LatePaymentFee      LatePaymentFee `json:"mora"`
	Payer               Payer          `json:"pagador"`
	Boleto              Boleto         `json:"boleto"`
	Pix                 Pix            `json:"pix"`
}

type Boleto struct {
	OurNumber string `json:"nossoNumero"`
	BarCode   string `json:"codigoBarras"`
	LineOne   string `json:"linhaDigitavel"`
}

type Pix struct {
	TxID            string `json:"txid"`
	PixCopyAndPaste string `json:"pixCopiaECola"`
}

// CreateCharge - Create a charge
func (i inter) CreateCharge(charge Charge) (*Charge, error) {
	var c Charge

	return &c, errors.New("not implemented")
}

// GetCharge - Get a charge
func (i inter) GetCharge(uuid string) (*Charge, error) {
	var c Charge

	return &c, errors.New("not implemented")
}

// DowloadCharge - Download a charge
func (i inter) DowloadCharge(uuid string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

// CancelCharge - Cancel a charge
func (i inter) CancelCharge(uuid string, reason string) error {
	return errors.New("not implemented")
}
