package models

import (
	"time"
)

type GetArticlesParams struct {
	Offset      *int  `form:"offset" json:"offset"  binding:"required"`
	Limit       *int  `form:"limit" json:"limit"  binding:"required"`
	Published   *bool `form:"published,default=true" json:"published"`
	Unpublished *bool `form:"unpublished,default=true" json:"unpublished"`
}

type GetLiveDataParams struct {
	Since *time.Time `form:"since" json:"since" binding:"required" time_format:"2006-01-02T15:04:05.000Z07:00"`
}

type SetLiveData struct {
	UplinkMessage SetUplinkMessage `json:"uplink_message"`
}

type SetUplinkMessage struct {
	DecodedPayload LiveData `json:"decoded_payload"`
}
