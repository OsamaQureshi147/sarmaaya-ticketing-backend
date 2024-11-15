package models

import (
	"gorm.io/gorm"
)

// ProductType defines the possible values for the product enum
type ProductType string

const (
	Zar    ProductType = "zar"
	Mahir  ProductType = "mahir"
	Dotcom ProductType = "dotcom"
)

type List struct {
	gorm.Model
	ListID             int64       `gorm:"column:list_id;not null" json:"list_id"`
	Product            ProductType `gorm:"type:product_type;not null" json:"product"`
	EmailFieldID       string      `gorm:"column:email_field_id;not null" json:"email_field_id"`
	AttachmentFieldID  string      `gorm:"column:attachment_field_id;not null" json:"attachment_field_id"`
	DescriptionFieldID string      `gorm:"column:description_field_id;not null" json:"description_field_id"`
}
