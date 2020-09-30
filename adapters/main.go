package adapters

import (
    "fmt"

    "github.com/faoztas/exrate/adapters/cbuae"
    "github.com/faoztas/exrate/adapters/tcmb"
    "github.com/faoztas/exrate/adapters/xe"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type AdapterInterface interface {
    LoadCredentials(adapter *models.Adapter)
    FetchExchangeRates(QuitChannel chan bool) error
    AddLogFields(fields logrus.Fields)
}

func Get(adapter string) (AdapterInterface, error) {
    switch adapter {
    case cbuae.Name:
        return new(cbuae.CBUAE), nil
    case tcmb.Name:
        return new(tcmb.TCMB), nil
    case xe.Name:
        return new(xe.XE), nil
    default:
        return nil, fmt.Errorf("%s: adapter not implemented", adapter)
    }
}
