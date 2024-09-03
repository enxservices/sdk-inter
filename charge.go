package intersdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	types2 "github.com/enxservices/sdk-inter/internal/types"
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

type QueryParamChargeList struct {
	InitialDate   *string `json:"dataInicial"`
	FinalDate     *string `json:"dataFinal"`
	FilterDataPor *string `json:"filtroDataPor"`
	Status        *string `json:"situacao"`
	Payer         *string `json:"pessoaPagadora"`
	Doc           *string `json:"cpfCnpjPessoaPagadora"`
	YourNumber    *string `json:"seuNumero"`
	CobType       *string `json:"tipoCobranca"`
	ItemPerPage   *int    `json:"paginacao.itensPorPagina"`
	CurrentPage   *int    `json:"paginacao.paginaAtual"`
	SortBy        *string `json:"ordernarPor"`
	TypeSort      *string `json:"tipoOrdenacao"`
}

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
	Fee          float64          `json:"taxa"`
	Code         CodeTypeDiscount `json:"codigo"`
	DaysQuantity int              `json:"quantidadeDias"`
}

type Fine struct {
	Fee   float64     `json:"taxa"`
	Code  CodeTypeFee `json:"codigo"`
	Value *float64    `json:"valor"`
}

type LatePaymentFee struct {
	Fee   float64             `json:"taxa"`
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

type CancelChargeRequest struct {
	Reason string `json:"motivoCancelamento"`
}

type CreateChargeResponse struct {
	SolicitationCode string `json:"codigoSolicitacao"`
}

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
	NominalValue        float64        `json:"valorNominal"`
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
}

type ChargeList struct {
	TotalPages       int  `json:"totalPaginas"`
	TotalElements    int  `json:"totalElementos"`
	SizePage         int  `json:"tamanhoPagina"`
	FistPage         bool `json:"primeiraPagina"`
	LastPage         bool `json:"ultimaPagina"`
	NumberOfElements int  `json:"numeroDeElementos"`
	Charges          []struct {
		Charge struct {
			SolicitationCode string       `json:"codigoSolicitacao"`
			YourNumber       string       `json:"seuNumero"`
			EmissionDate     string       `json:"dataEmissao"`
			NominalValue     string       `json:"valorNominal"`
			DueDate          string       `json:"dataVencimento"`
			ChargeType       ChargeType   `json:"tipoCobranca"`
			Status           ChargeStatus `json:"situacao"`
			StatusDate       string       `json:"dataSituacao"`
			Payer            struct {
				Name string `json:"nome"`
				Doc  string `json:"cpfCnpj"`
			} `json:"pagador"`
		} `json:"cobranca"`
		Boleto Boleto `json:"boleto"`
		Pix    Pix    `json:"pix"`
	} `json:"cobrancas"`
}

type ChargeResponse struct {
	Charge Charge `json:"cobranca"`
	Boleto Boleto `json:"boleto"`
	Pix    Pix    `json:"pix"`
}

type ErrorCancelCharge struct {
	Status     int         `json:"status"`
	Title      string      `json:"title"`
	Detail     string      `json:"detail"`
	Violations []Violation `json:"violacoes"`
}

type Violation struct {
	Reason   string  `json:"razao"`
	Property string  `json:"propriedade"`
	Value    *string `json:"valor"`
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
func (i inter) CreateCharge(charge CreateChargeRequest) (string, error) {
	dueDate, err := time.Parse("2006-01-02", charge.DueDate.Format("2006-01-02"))
	if err != nil {
		return "", err
	}

	charge.DueDate = dueDate

	payload, err := json.Marshal(charge)
	if err != nil {
		return "", err
	}

	token, err := i.Oauth.GetAccessToken("boleto-cobranca.write")
	if err != nil {
		return "", err
	}

	res, err := sendRequest(i.client, "POST", fmt.Sprintf("%s/%s", i.BaseURL, types2.CobPixBoletoUrl), token, payload)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		fmt.Println(string(resBody))
		return "", errors.New(string(resBody))
	}

	var solicitationCode CreateChargeResponse
	if err := json.Unmarshal(resBody, &solicitationCode); err != nil {
		return "", err
	}

	return solicitationCode.SolicitationCode, nil
}

// GetCharge - Get a charge
func (i inter) GetCharge(solicitationCode string) (*ChargeResponse, error) {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.read")
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(i.client, "GET", fmt.Sprintf("%s/%s/%s", i.BaseURL, types2.CobPixBoletoUrl, solicitationCode), token, nil)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		fmt.Println(string(resBody))
		return &ChargeResponse{}, errors.New(string(resBody))
	}

	var charge ChargeResponse
	if err := json.Unmarshal(resBody, &charge); err != nil {
		return nil, err
	}

	return &charge, nil
}

// DowloadCharge - Download a charge
func (i inter) DowloadCharge(solicitationCode string) (string, error) {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.read")
	if err != nil {
		return "", err
	}

	type Response struct {
		Pdf string `json:"pdf"`
	}

	res, err := sendRequest(i.client, "GET", fmt.Sprintf("%s/%s/%s/pdf", i.BaseURL, types2.CobPixBoletoUrl, solicitationCode), token, nil)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		fmt.Println(string(resBody))
		return "", errors.New(string(resBody))
	}

	var pdf Response
	if err := json.Unmarshal(resBody, &pdf); err != nil {
		return "", err
	}

	return pdf.Pdf, nil
}

func buildChargeListURL(baseURL string, params QueryParamChargeList) string {
	url, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	query := url.Query()

	if params.InitialDate != nil {
		query.Add("dataInicial", *params.InitialDate)
	}
	if params.FinalDate != nil {
		query.Add("dataFinal", *params.FinalDate)
	}
	if params.FilterDataPor != nil {
		query.Add("filtroDataPor", *params.FilterDataPor)
	}
	if params.Status != nil {
		query.Add("situacao", *params.Status)
	}
	if params.Payer != nil {
		query.Add("pessoaPagadora", *params.Payer)
	}
	if params.Doc != nil {
		query.Add("cpfCnpjPessoaPagadora", *params.Doc)
	}
	if params.YourNumber != nil {
		query.Add("seuNumero", *params.YourNumber)
	}
	if params.CobType != nil {
		query.Add("tipoCobranca", *params.CobType)
	}
	if params.ItemPerPage != nil {
		query.Add("paginacao.itensPorPagina", strconv.Itoa(*params.ItemPerPage))
	}
	if params.CurrentPage != nil {
		query.Add("paginacao.paginaAtual", strconv.Itoa(*params.CurrentPage))
	}
	if params.SortBy != nil {
		query.Add("ordernarPor", *params.SortBy)
	}
	if params.TypeSort != nil {
		query.Add("tipoOrdenacao", *params.TypeSort)
	}

	url.RawQuery = query.Encode()
	return url.String()
}

func (i inter) GetChargeList(params QueryParamChargeList) (*ChargeList, error) {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.read")
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(i.client, "GET", buildChargeListURL(fmt.Sprintf("%s/%s", i.BaseURL, types2.CobPixBoletoUrl), params), token, nil)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}

	var chargeList ChargeList
	if err := json.Unmarshal(resBody, &chargeList); err != nil {
		return nil, err
	}

	return &chargeList, nil
}

// CancelCharge - Cancel a charge
func (i inter) CancelCharge(solicitationCode string, reason string) error {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.write")
	if err != nil {
		return err
	}

	payload, err := json.Marshal(CancelChargeRequest{Reason: reason})

	if err != nil {
		return err
	}

	res, err := sendRequest(i.client, "POST", fmt.Sprintf("%s/%s/%s/cancelar", i.BaseURL, types2.CobPixBoletoUrl, solicitationCode), token, payload)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 202 {
		var errorCancelCharge ErrorCancelCharge
		if err := json.NewDecoder(res.Body).Decode(&errorCancelCharge); err != nil {
			return err
		}
		return errors.New("error canceling charge")
	}

	return nil
}
