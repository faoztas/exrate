package models

import (
    "time"

    "github.com/shopspring/decimal"
)

type ExchangeRule struct {
    ID               uint            `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt        *time.Time      `json:"created_at,omitempty" gorm:""`
    UpdatedAt        *time.Time      `json:"updated_at,omitempty" gorm:""`
    SourceCurrency   Currency        `json:"source_currency" gorm:"foreignkey:SourceCurrencyID;association_autoupdate:false;association_autocreate:false"`
    SourceCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    TargetCurrency   Currency        `json:"target_currency" gorm:"foreignkey:TargetCurrencyID;association_autoupdate:false;association_autocreate:false"`
    TargetCurrencyID string          `json:"-" gorm:"type:varchar(3);not null"`
    BaseCurrency     Currency        `json:"base_currency" gorm:"foreignkey:BaseCurrencyID;association_autoupdate:false;association_autocreate:false"`
    BaseCurrencyID   string          `json:"-" gorm:"type:varchar(3);not null"`
    CommissionRate   decimal.Decimal `json:"commission_rate" gorm:"type:decimal(3,2);not null"`
    UpdateFrequency  uint            `json:"update_frequency" gorm:"type:numeric;not null"`
    IsActive         bool            `json:"is_active" gorm:"type:boolean"`
    IsFixed          bool            `json:"is_fixed" gorm:"type:boolean"`
}
