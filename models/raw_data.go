package models

import (
    "time"
)
type RawData struct {
    ID               int        `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt        *time.Time `json:"created_at,omitempty" gorm:""`
    UpdatedAt        *time.Time `json:"updated_at,omitempty" gorm:""`
    SourceCurrency   Currency   `json:"source_currency" gorm:"foreignkey:ID;association_foreignkey:SourceCurrencyID"`
    SourceCurrencyID string     `json:"-" gorm:"type:varchar(3);not null"`
    TargetCurrency   Currency   `json:"target_currency" gorm:"foreignkey:ID;association_foreignkey:TargetCurrencyID"`
    TargetCurrencyID string     `json:"-" gorm:"type:varchar(3);not null"`
    RawRequest       string     `json:"raw_request,omitempty" gorm:""`
    RawResponse      string     `json:"raw_response,omitempty" gorm:""`
    IsSuccess        bool       `json:"is_success" gorm:"type:boolean"`
}
