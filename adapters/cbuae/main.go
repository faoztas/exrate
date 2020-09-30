package cbuae

import (
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type CBUAE struct {
    Adapter *models.Adapter
    logger  *logrus.Entry
}

func (a CBUAE) LoadCredentials(adapter *models.Adapter) {
    panic("implement me")
}

func (a CBUAE) FetchExchangeRates(QuitChannel chan bool) error {
    panic("implement me")
}

func (a CBUAE) AddLogFields(fields logrus.Fields) {
    panic("implement me")
}
