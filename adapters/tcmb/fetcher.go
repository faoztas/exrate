package tcmb

import (
    "encoding/xml"
    "fmt"
    "net/http"
    "time"

    "github.com/faoztas/exrate/db"
    "github.com/faoztas/exrate/models"
    "github.com/faoztas/exrate/utils"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

func (a *TCMB) FetchExchangeRates() {
    request := utils.R{
        Method:  http.MethodGet,
        Timeout: &a.Credential.Timeout,
        URL:     fmt.Sprintf(a.Credential.BaseURL, today),
        Body:    nil,
        Header:  nil,
    }
    requestTime := time.Now().UTC()
    response, body, err := utils.Request(request)
    responseTime := time.Now().UTC()

    serviceLog := models.ServiceLog{
        UUID:           uuid.New().String(),
        AdapterID:      a.Adapter.ID,
        RequestBody:    utils.ToJson(request.Body),
        ResponseBody:   utils.Xml2Json(body),
        ExternalData:   utils.ToJson(request.Header),
        LogType:        utils.AdapterTCMB,
        RequestDate:    requestTime,
        ResponseDate:   responseTime,
        HTTPStatus:     utils.ToHttpStatusCode(response),
        HTTPMethod:     request.Method,
        DestinationURL: request.URL,
        IsSuccess:      false,
    }
    if err != nil {
        db.DB.Save(&serviceLog)
        a.logger.Errorf("%v Request: %+v Response: %+v", err, request, response)
        return
    }
    serviceLog.IsSuccess = true
    db.DB.Save(&serviceLog)

    rowData := new(Binding)
    err = xml.Unmarshal(body, &rowData)
    if err != nil {
        a.logger.Errorf("Unmarshalling Error: %v ", err)
        return
    }

    for k := range rowData.Currency {
        if rowData.Currency[k].CurrencyCode == "XDR" {
            continue
        }

        forexBuying, err := decimal.NewFromString(rowData.Currency[k].ForexBuying)
        if err != nil {
            a.logger.Error(err)
            continue
        }

        forexSelling, err := decimal.NewFromString(rowData.Currency[k].ForexSelling)
        if err != nil {
            a.logger.Error(err)
            continue
        }

        if rowData.Currency[k].Unit != "1" {
            unit, err := decimal.NewFromString(rowData.Currency[k].Unit)
            if err != nil {
                a.logger.Error(err)
                continue
            }
            forexBuying = forexBuying.Div(unit)
            forexSelling = forexSelling.Div(unit)
        }
        rate := forexBuying.Add(forexSelling).Div(decimal.NewFromInt(2))

        if err := db.DB.Create(&models.ExchangeHistory{
            SourceCurrencyID: rowData.Currency[k].CurrencyCode,
            TargetCurrencyID: a.targetCurrency,
            CurrencyRate:     utils.DecimalToFixed(rate, 13),
            ServiceLogID:     serviceLog.ID,
        }).Error; err != nil {
            a.logger.Errorf("Insert Error: %v ", err)
            continue
        }
    }
}
