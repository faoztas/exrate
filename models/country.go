package models

import (
    "time"
)

type Country struct {
    ID          string    `json:"id" gorm:"type:varchar(2);primary_key"`
    CreatedAt   time.Time `json:"created_at" gorm:""`
    UpdatedAt   time.Time `json:"updated_at" gorm:""`
    CreatedBy   User      `json:"created_by,omitempty" gorm:"foreignkey:CreatedByID"`
    CreatedByID int       `json:"created_by_id" gorm:"type:integer"`
    Version     int       `json:"version" gorm:"type:integer"`
    IsActive    bool      `json:"is_active" gorm:"type:boolean"`
    Iso3Code    string    `json:"iso3_code" gorm:"type:varchar(3)"`
    CountryName string    `json:"country_name" gorm:"type:varchar(45)"`
    Currency    Currency  `json:"currency,omitempty" gorm:"foreignkey:CurrencyID"`
    CurrencyID  string    `json:"currency_id" gorm:"type:varchar(3)" sql:"default: null"`
    NumCode     int       `json:"num_code" gorm:"type:integer"`
    PhoneCode   int       `json:"phone_code" gorm:"type:integer"`
}
