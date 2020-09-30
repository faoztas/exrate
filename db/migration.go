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
        &models.ExchangeHistory{},
        &models.ExchangeRule{},
        &models.RawData{},
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