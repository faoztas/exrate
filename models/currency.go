package models

import (
    "time"
)

type Currency struct {
    ID             string     `json:"id" gorm:"column:id;PRIMARY_KEY"`
    CreatedAt      *time.Time `json:"created_at,omitempty" gorm:""`
    UpdatedAt      *time.Time `json:"updated_at,omitempty" gorm:""`
    CreatedBy      *User      `json:"created_by,omitempty" gorm:"foreignkey:CreatedByID"`
    CreatedByID    int        `json:"-" gorm:"type:int"`
    Version        int        `json:"version" gorm:""`
    IsActive       bool       `json:"is_active" gorm:"type:boolean"`
    HasActiveFlow  bool       `json:"has_active_flow" gorm:"type:boolean"`
    CurrencyName   string     `json:"currency_name" gorm:"type:varchar(45)"`
    CurrencyNameTr string     `json:"currency_name_tr" gorm:"type:varchar(60)"`
    Symbol         string     `json:"symbol" gorm:"type:varchar(45)"`
}