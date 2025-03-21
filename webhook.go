package intersdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	types2 "github.com/enxservices/sdk-inter/internal/types"
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

func (i inter) CreateWebhook(webhookUrl string) error {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.write")
	if err != nil {
		return err
	}

	payload := CreateWebhook{
		WebhookUrl: webhookUrl,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	res, err := sendRequest(i.client, http.MethodPut, fmt.Sprintf("%s/%s", i.BaseURL, types2.CobWebHookUrl), token, jsonData)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("falha ao criar webhook: %+v", res)
	}

	return nil
}

func (i inter) GetWebhook() (*Webhook, error) {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.read")
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(i.client, http.MethodGet, fmt.Sprintf("%s/%s", i.BaseURL, types2.CobWebHookUrl), token, nil)
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

func (i inter) DeleteWebhook() (*WebhookError, error) {
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.write")
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(i.client, http.MethodGet, types2.CobWebHookUrl, token, nil)
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
	token, err := i.Oauth.GetAccessToken("boleto-cobranca.read")
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(types2.CobWebHookUrlCallbacks)
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

	res, err := sendRequest(i.client, http.MethodGet, baseUrl.String(), token, nil)
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
