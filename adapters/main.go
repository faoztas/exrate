package adapters

import (
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type AdapterInterface interface {
    LoadCredentials(adapter *models.Adapter, credential *config.Credential)
    FetchExchangeRates() error
    AddLogFields(fields logrus.Fields)
}
