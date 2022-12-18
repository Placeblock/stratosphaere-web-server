package models

type GetArticlesParams struct {
	Offset      *int  `form:"offset" json:"offset"  binding:"required"`
	Limit       *int  `form:"limit" json:"limit"  binding:"required"`
	Published   *bool `form:"published,default=true" json:"published"`
	Unpublished *bool `form:"unpublished,default=true" json:"unpublished"`
}
