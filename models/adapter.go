package models

import (
    "time"

    "gorm.io/datatypes"
)

type Adapter struct {
    ID            int            `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt     time.Time      `json:"created_at" gorm:""`
    UpdatedAt     time.Time      `json:"updated_at" gorm:""`
    CreatedBy     User           `json:"created_by,omitempty" gorm:"foreignkey:CreatedByID"`
    CreatedByID   int            `json:"created_by_id" gorm:"type:integer"`
    Name          string         `json:"name" gorm:"type:varchar(50)"`
    Configuration datatypes.JSON `json:"configuration,omitempty" gorm:""`
    IsActive      bool           `json:"is_active" gorm:"type:boolean"`
}
