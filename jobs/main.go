package job

import (
    "github.com/faoztas/exrate/adapters"
    "github.com/faoztas/exrate/config"
    "github.com/sirupsen/logrus"
)

func Jobs() {
    logrus.Info("Starting Scheduling jobs")

    if config.Env.FeatureFlags.Flag {
        for _, a := range GetAdapter() {

            adapter, err := adapters.Use(a.Name)
            if err != nil {
                logrus.Error(err)
                continue
            }
            credential := config.Env.Adapters[a.Name]

            adapter.LoadCredentials(&a, &credential)
            adapter.AddLogFields(logrus.Fields{
                "module": "adapters",
            })

            err = adapter.FetchExchangeRates()
            logrus.Error(err)

            // RunRoutineForever(adapter.FetchExchangeRates, 1*time.Minute)
        }
    } else {
        logrus.Info("Flag not enabled")
    }
}
