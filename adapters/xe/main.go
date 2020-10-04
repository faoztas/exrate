package xe

import (
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type XE struct {
    Adapter *models.Adapter
    logger  *logrus.Entry
}

func (a XE) LoadCredentials(adapter *models.Adapter, credential *config.Credential) {
    panic("implement me")
}

func (a XE) FetchExchangeRates() error {
    panic("implement me")
}

func (a XE) AddLogFields(fields logrus.Fields) {
    panic("implement me")
}
