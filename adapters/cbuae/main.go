package cbuae

import (
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type CBUAE struct {
    Adapter *models.Adapter
    logger  *logrus.Entry
}

func (a CBUAE) LoadCredentials(adapter *models.Adapter, credential *config.Credential) {
    panic("implement me")
}

func (a CBUAE) FetchExchangeRates() error {
    panic("implement me")
}

func (a CBUAE) AddLogFields(fields logrus.Fields) {
    panic("implement me")
}
