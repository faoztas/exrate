package xe

import (
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type XE struct {
    Adapter *models.Adapter
    logger  *logrus.Entry
}

func (a XE) LoadCredentials(adapter *models.Adapter) {
    panic("implement me")
}

func (a XE) FetchExchangeRates(QuitChannel chan bool) error {
    panic("implement me")
}

func (a XE) AddLogFields(fields logrus.Fields) {
    panic("implement me")
}
