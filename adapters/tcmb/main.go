package tcmb

import (
    "github.com/faoztas/exrate/config"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

type TCMB struct {
    Adapter        *models.Adapter
    Credential     *config.Credential
    targetCurrency string
    logger         *logrus.Entry
}

func (a TCMB) LoadCredentials(adapter *models.Adapter, credential *config.Credential) {
    a.Adapter = adapter
    a.Credential = credential
    a.targetCurrency = "TRY"
    a.logger = logrus.WithFields(logrus.Fields{
        "adapter": adapter.Name,
    })
}

func (a TCMB) FetchExchangeRates() error {
    if err := a.fetcher(); err != nil {
        return err
    }
    return nil
}

func (a TCMB) AddLogFields(fields logrus.Fields) {
    a.logger = a.logger.WithFields(fields)
}
