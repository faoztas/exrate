package models

import (
    "time"
)

type Role struct {
    ID        uint       `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    Name      string     `json:"name" gorm:"type:varchar;unique_index;not null"`
    Status    bool       `json:"status" gorm:"type:boolean"`
}