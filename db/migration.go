package db

import (
    "github.com/faoztas/exrate/models"
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
            panic(err)
        }
    }
}