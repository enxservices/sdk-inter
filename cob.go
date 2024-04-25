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
	valueFixedInfoData CodeTypeDiscount = "VALORFIXODATAINFORMADA"
	percentageInfoData CodeTypeDiscount = "PERCENTUALDATAINFORMADA"
)

type CodeTypeFee string

const (
	feePercentage CodeTypeFee = "PERCENTUAL"
	feeFix        CodeTypeFee = "VALORFIXO"
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
	email      string     `json:"email"`
	ddd        string     `json:"ddd"`
	phone      string     `json:"telefone"`
	number     string     `json:"numero"`
	complement string     `json:"complemento"`
	doc        string     `json:"cpfCnpj"`
	kindPeople KindPeople `json:"tipoPessoa"`
	name       string     `json:"nome"`
	address    string     `json:"endereco"`
	district   string     `json:"bairro"`
	city       string     `json:"cidade"`
	state      State      `json:"uf"`
	zipCode    string     `json:"cep"`
}

type Message struct {
	lineOne   *string `json:"linha1"`
	lineTwo   *string `json:"linha2"`
	lineThree *string `json:"linha3"`
	lineFour  *string `json:"linha4"`
	lineFive  *string `json:"linha5"`
}

type Discount struct {
	fee          int              `json:"taxa"`
	code         CodeTypeDiscount `json:"codigo"`
	daysQuantity int              `json:"quantidadeDias"`
}

type Fine struct {
	fee  int         `json:"taxa"`
	code CodeTypeFee `json:"codigo"`
}

type LatePaymentFee struct {
	fee  int                 `json:"taxa"`
	code CodeTypeLatePayment `json:"codigo"`
}

type Beneficiary struct {
	doc        string     `json:"cpfCnpj"`
	kindPeople KindPeople `json:"tipoPessoa"`
	name       string     `json:"nome"`
	address    string     `json:"endereco"`
	district   string     `json:"bairro"`
	city       string     `json:"cidade"`
	state      State      `json:"uf"`
	zipCode    string     `json:"cep"`
}

// Date Format YYYY-MM-DD
type Charge struct {
	yourNumber       string          `json:"seuNumero"`
	nominalValue     float64         `json:"valorNominal"`
	dueDate          time.Time       `json:"dataVencimento"`
	daysAfterDue     int             `json:"numDiasAgenda"`
	payer            Payer           `json:"pagador"`
	discount         *Discount       `json:"desconto,omitempty"`
	fine             *Fine           `json:"multa,omitempty"`
	latePayment      *LatePaymentFee `json:"juros,omitempty"`
	message          *Message        `json:"mensagem,omitempty"`
	finalBeneficiary Beneficiary     `json:"beneficiarioFinal"`
}

//Type constructors

func NewBeneficiary(doc, name, address, district, city, zipCode string, kindPeople KindPeople, state State) Beneficiary {
	return Beneficiary{
		doc:        doc,
		kindPeople: kindPeople,
		name:       name,
		address:    address,
		district:   district,
		city:       city,
		state:      state,
		zipCode:    zipCode,
	}
}

func NewPayer(email, ddd, phone, number, complement, doc, name, address, district, city, zipCode string, kindPeople KindPeople, state State) Payer {
	return Payer{
		email:      email,
		ddd:        ddd,
		phone:      phone,
		number:     number,
		complement: complement,
		doc:        doc,
		kindPeople: kindPeople,
		name:       name,
		address:    address,
		district:   district,
		city:       city,
		state:      state,
		zipCode:    zipCode,
	}
}

func NewDiscount(fee, daysQuantity int, code CodeTypeDiscount) Discount {
	return Discount{
		fee:          fee,
		daysQuantity: daysQuantity,
		code:         code,
	}
}

func NewFine(fee int, code CodeTypeFee) Fine {
	return Fine{
		fee:  fee,
		code: code,
	}
}

func NewMessage(lineOne, lineTwo, lineThree, lineFour, lineFive *string) Message {
	return Message{
		lineOne:   lineOne,
		lineTwo:   lineTwo,
		lineThree: lineThree,
		lineFour:  lineFour,
		lineFive:  lineFive,
	}
}

func NewLatePaymentFee(fee int, code CodeTypeLatePayment) LatePaymentFee {
	return LatePaymentFee{
		fee:  fee,
		code: code,
	}
}

func NewCharge(yourNumber string, nominalValue float64, dueDate time.Time, daysAfterDue int, payer Payer, discount *Discount, latePayment *LatePaymentFee, fine *Fine, message *Message, finalBeneficiary Beneficiary) Charge {
	return Charge{
		yourNumber:       yourNumber,
		nominalValue:     nominalValue,
		dueDate:          dueDate,
		daysAfterDue:     daysAfterDue,
		payer:            payer,
		discount:         discount,
		latePayment:      latePayment,
		fine:             fine,
		message:          message,
		finalBeneficiary: finalBeneficiary,
	}
}
