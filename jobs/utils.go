package job

import (
    "errors"
    "runtime/debug"
    "time"

    "github.com/faoztas/exrate/db"
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
)

func RunRoutineForever(fn func(), delay time.Duration) {
    go func() {
        defer retryRecover(fn, delay)
        fn()
    }()
}

func retryRecover(fn func(), delay time.Duration) {
    if r := recover(); r != nil {
        var err error
        switch x := r.(type) {
        case string:
            err = errors.New(x)
        case error:
            err = x
        default:
            err = errors.New("unknown panic")
        }

        logrus.WithFields(logrus.Fields{
            "trace": string(debug.Stack()),
        }).WithError(err).Error("RunRoutineForever recovered")
    }

    logrus.Infof("Task will be run after %v seconds", delay.Seconds())
    time.Sleep(delay)
    RunRoutineForever(fn, delay)
}

func GetAdapter() []models.Adapter {
     var adapters []models.Adapter
    if err := db.DB.Where("is_active = ?", true).Find(&adapters).Error; err != nil {
        logrus.Error(err)
    }
    return adapters
}

