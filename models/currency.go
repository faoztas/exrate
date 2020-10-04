package models

import (
    "time"
)

type Currency struct {
    ID            string    `json:"id" gorm:"type:varchar(3);column:id;PRIMARY_KEY"`
    CreatedAt     time.Time `json:"created_at" gorm:""`
    UpdatedAt     time.Time `json:"updated_at" gorm:""`
    CreatedBy     User      `json:"created_by,omitempty" gorm:"foreignkey:CreatedByID;association_autoupdate:false;association_autocreate:false"`
    CreatedByID   int       `json:"-" gorm:"type:integer"`
    IsActive      bool      `json:"is_active" gorm:"type:boolean"`
    CurrencyName  string    `json:"currency_name" gorm:"type:varchar(45)"`
    Symbol        string    `json:"symbol" gorm:"type:varchar(45)"`
    NewCurrencyID string    `json:"new_currency_id" gorm:"type:varchar(3)"`
}
