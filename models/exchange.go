package models

import (
    "time"

    "github.com/shopspring/decimal"
)

type Exchange struct {
    ID               int             `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt        time.Time       `json:"created_at" gorm:""`
    UpdatedAt        time.Time       `json:"updated_at" gorm:""`
    SourceCurrency   Currency        `json:"source_currency" gorm:"foreignkey:SourceCurrencyID;association_autoupdate:false;association_autocreate:false"`
    SourceCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    TargetCurrency   Currency        `json:"target_currency" gorm:"foreignkey:TargetCurrencyID;association_autoupdate:false;association_autocreate:false"`
    TargetCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    CurrencyRate     decimal.Decimal `json:"currency_rate" gorm:"type:decimal(25,13);not null"`
    ExchangeRateTime *time.Time      `json:"exchange_rate_time" gorm:""`
}
