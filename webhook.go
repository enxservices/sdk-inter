package intersdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/enxservices/sdk-inter/internal/types"
)

type Notification struct {
	WebHookUrl       string           `json:"webhookUrl"`
	Payload          []WebhookPayload `json:"payload"`
	Tries            int              `json:"numeroTentativa"`
	NotificationTime time.Time        `json:"dataHoraDisparo"`
	Send             bool             `json:"sucesso"`
	HttpStatusCode   int              `json:"httpStatus"`
	ErrorMessage     string           `json:"mensagemErro"`
}

type CreateWebhook struct {
	WebhookUrl string `json:"webhookUrl"`
}

type WebhookPayload struct {
	RequestCode         string       `json:"codigoSolicitacao"`
	YourNumber          string       `json:"seuNumero"`
	Status              ChargeStatus `json:"situacao"`
	StatusDateTime      string       `json:"dataHoraSituacao"`
	TotalReceivedAmount *string      `json:"valorTotalRecebido,omitempty"`
	ReceivingSource     *string      `json:"origemRecebimento,omitempty"`
	OurNumber           string       `json:"nossoNumero"`
	Barcode             string       `json:"codigoBarras"`
	DigitableLine       string       `json:"linhaDigitavel"`
	TxID                string       `json:"txid"`
	PixCopyPaste        string       `json:"pixCopiaECola"`
}

type Webhook struct {
	WebhookUrl string    `json:"webhookUrl"`
	CreatedAt  time.Time `json:"criacao"`
	UpdatedAt  time.Time `json:"atualizacao"`
}

type WebhookError struct {
	Title      string       `json:"title"`
	StatusCode int          `json:"status"`
	Detail     string       `json:"detail"`
	Violations []Violations `json:"violacoes"`
}

func (e WebhookError) String() string {
	var violationsStr strings.Builder
	for _, v := range e.Violations {
		violationsStr.WriteString(fmt.Sprintf("Reason: %s, Property: %s, Value: %s\n", v.Reason, v.Property, v.Value))
	}
	return fmt.Sprintf(
		"Title: %s\nStatusCode: %d\nDetail: %s\nViolations:\n%s",
		e.Title, e.StatusCode, e.Detail, violationsStr.String(),
	)
}

type Violations struct {
	Reason   string `json:"razao"`
	Property string `json:"propriedade"`
	Value    string `json:"valor"`
}

/**
* SizePage [10...50] Default 20
* SolicitaonCode <uuid>
 */
type Queries struct {
	StartDate        time.Time
	EndDate          time.Time
	Page             *int
	SizePage         *int
	SolicitacionCode *string
}

type CallBackSend struct {
	TotalItems int            `json:"totalElementos"`
	TotalPages int            `json:"totalPaginas"`
	FirstPage  bool           `json:"primeiraPagina"`
	LastPage   bool           `json:"ultimaPagina"`
	Data       []Notification `json:"data"`
}


func (i inter) Create(webhookUrl string) error {
	token := i.Oauth.GetAccessToken(types.Scope("boleto-cobranca.write"))

	payload := CreateWebhook{
		WebhookUrl: webhookUrl,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	res, err := sendRequest(i.client, "PUT", types.CobWebHookUrl, token, jsonData)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return errors.New("unable to create webhook")
	}

	return nil
}


func (i inter) Get() (*Webhook, error) {
	token := i.Oauth.GetAccessToken(types.Scope("boleto-cobranca.read"))

	res, err := sendRequest(i.client, "GET", types.CobWebHookUrl, token, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("no webhooks exists")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var webhook Webhook
	if err := json.Unmarshal(resBody, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (i inter) Delete() (*WebhookError, error) {
	token := i.Oauth.GetAccessToken(types.Scope("boleto-cobranca.write"))

	res, err := sendRequest(i.client, "GET", types.CobWebHookUrl, token, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var WebhookError WebhookError
		if err := json.Unmarshal(resBody, &WebhookError); err != nil {
			return nil, err
		}

		return &WebhookError, errors.New("could not delete webhook")
	}

	return nil, nil
}

func (i inter) GetAllCallbackSend(queries Queries) (*CallBackSend, error) {
	token := i.Oauth.GetAccessToken(types.Scope("boleto-cobranca.read"))

	baseUrl, err := url.Parse(types.CobWebHookUrlCallbacks)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	layout := "2006-01-02T15:04:05.000Z07:00"

	params.Add("start_date", queries.StartDate.Format(layout))
	params.Add("end_date", queries.EndDate.Format(layout))

	if queries.Page != nil {
		params.Add("page", fmt.Sprintf("%d", *queries.Page))
	}
	if queries.SizePage != nil {
		params.Add("size_page", fmt.Sprintf("%d", *queries.SizePage))
	}
	if queries.SolicitacionCode != nil {
		params.Add("solicitacion_code", *queries.SolicitacionCode)
	}

	baseUrl.RawQuery = params.Encode()

	res, err := sendRequest(i.client, "GET", baseUrl.String(), token, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		var webhookUrlError WebhookError
		if err := json.Unmarshal(resBody, &webhookUrlError); err != nil {
			return nil, err
		}
		return nil, errors.New(webhookUrlError.String())
	}

	var callbackSend CallBackSend
	if err := json.Unmarshal(resBody, &callbackSend); err != nil {
		return nil, err
	}

	return &callbackSend, nil
}
