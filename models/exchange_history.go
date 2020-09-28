package models

import (
    "time"

    "github.com/shopspring/decimal"
)

type ExchangeHistory struct {
    ID               int             `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt        *time.Time      `json:"created_at,omitempty" gorm:""`
    UpdatedAt        *time.Time      `json:"updated_at,omitempty" gorm:""`
    SourceCurrency   Currency        `json:"source_currency" gorm:"foreignkey:ID;association_foreignkey:SourceCurrencyID"`
    SourceCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    TargetCurrency   Currency        `json:"target_currency" gorm:"foreignkey:ID;association_foreignkey:TargetCurrencyID"`
    TargetCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    CurrencyRate     decimal.Decimal `json:"currency_rate" gorm:"type:decimal(19,8);not null"`
    RawData          RawData         `json:"raw_data" gorm:"foreignkey:ID;association_foreignkey:RawDataID"`
    RawDataID        string          `json:"-" gorm:"type:int;not null"`
}
