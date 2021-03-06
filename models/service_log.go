package models

import (
    "time"

    "gorm.io/datatypes"
)

type ServiceLog struct {
    ID             int            `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt      time.Time      `json:"created_at,omitempty" gorm:""`
    UUID           string         `json:"uuid" gorm:"type:varchar(36);"`
    Adapter        Adapter        `json:"adapter,omitempty" gorm:"foreignkey:AdapterID;association_autoupdate:false;association_autocreate:false"`
    AdapterID      int            `json:"adapter_id" gorm:"type:integer"`
    RequestBody    datatypes.JSON `json:"request_body,omitempty" gorm:""`
    ResponseBody   datatypes.JSON `json:"response_body,omitempty" gorm:""`
    ExternalData   datatypes.JSON `json:"external_data,omitempty" gorm:""`
    LogType        string         `json:"log_type" gorm:"type:varchar(50)"`
    RequestDate    time.Time      `json:"request_date" gorm:""`
    ResponseDate   time.Time      `json:"response_date" gorm:""`
    HTTPStatus     int            `json:"http_status" gorm:"type:integer"`
    HTTPMethod     string         `json:"http_method" gorm:"type:varchar(10)"`
    DestinationURL string         `json:"destination_url" gorm:"type:varchar(250)"`
    IsSuccess      bool           `json:"is_success" gorm:"type:boolean"`
}
