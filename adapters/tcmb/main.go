package tcmb

import (
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type TCMB struct {
    Adapter *models.Adapter
    logger  *logrus.Entry
}

func (a TCMB) LoadCredentials(adapter *models.Adapter) {
    panic("implement me")
}

func (a TCMB) FetchExchangeRates(QuitChannel chan bool) error {
    panic("implement me")
}

func (a TCMB) AddLogFields(fields logrus.Fields) {
    panic("implement me")
}
