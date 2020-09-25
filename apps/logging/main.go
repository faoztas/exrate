package logging

import (
    "time"

    "github.com/sirupsen/logrus"

    "github.com/faoztas/exrate/config"
)

func InitLogging() {
    logLevelMap := map[string]logrus.Level{
        "fatal":   logrus.FatalLevel,
        "warning": logrus.WarnLevel,
        "info":    logrus.InfoLevel,
        "debug":   logrus.DebugLevel,
        "trace":   logrus.TraceLevel,
        "error":   logrus.ErrorLevel,
    }

    // default level is info
    if v, ok := logLevelMap[config.Env.Logging.Level]; ok {
        logrus.SetLevel(v)
    }

    if config.Env.Logging.Format == "json" {
        logrus.SetFormatter(&logrus.JSONFormatter{})
    } else if config.Env.Logging.Format == "text" {
        logrus.SetFormatter(&logrus.TextFormatter{})
    }

    // For catching panic or fatal errors, we are waiting before exist.
    logrus.RegisterExitHandler(func() {
        time.Sleep(3 * time.Second)
    })
}
