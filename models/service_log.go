package models

import (
    "time"

    "gorm.io/datatypes"
)

type ServiceLog struct {
    ID             uint           `json:"id" gorm:"type:serial;primary_key;AUTO_INCREMENT"`
    CreatedAt      *time.Time     `json:"created_at,omitempty" gorm:""`
    RequestBody    datatypes.JSON `json:"request_body,omitempty" gorm:""`
    ResponseBody   datatypes.JSON `json:"response_body,omitempty" gorm:""`
    ExternalData   datatypes.JSON `json:"external_data,omitempty" gorm:""`
    LogType        string         `json:"log_type" gorm:""`
    RequestDate    *time.Time     `json:"request_date" gorm:""`
    ResponseDate   *time.Time     `json:"response_date" gorm:""`
    HTTPStatus     uint           `json:"http_status" gorm:""`
    HTTPMethod     string         `json:"http_method" gorm:""`
    DestinationURL string         `json:"destination_url" gorm:""`
    IsSuccess      bool           `json:"is_success" gorm:"type:boolean"`
}
