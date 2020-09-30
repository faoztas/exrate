package models

import (
    "time"
)

type User struct {
    ID        uint       `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    Username  string     `json:"user_name" gorm:"type:varchar(50);"`
    Email     string     `json:"email" gorm:"type:varchar(50);"`
    Role      *Role      `json:"role,omitempty" gorm:"foreignkey:RoleID"`
    RoleID    uint       `json:"-" gorm:"type:int;not null"`
    Status    bool       `json:"status"`
    FirstName string     `json:"first_name" gorm:"type:varchar(50);"`
    LastName  string     `json:"last_name" gorm:"type:varchar(50);"`
    Password  string     `json:"password" gorm:"type:varchar(100);"`
}
