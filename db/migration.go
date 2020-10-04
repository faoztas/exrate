package db

import (
    "github.com/faoztas/exrate/models"
    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    objects := []interface{}{
        &models.User{},
        &models.Role{},
        &models.Currency{},
        &models.Country{},
        &models.Exchange{},
        &models.ServiceLog{},
        &models.ExchangeHistory{},
        &models.ExchangeRule{},
    }

    for _, item := range objects {
        if item == nil {
            continue
        }
        if err := db.AutoMigrate(item); err != nil {
            logrus.Error(err)
            panic(err)
        }
    }

    if err := CustomMigrations(); err != nil {
        logrus.Error(err)
    }
}
