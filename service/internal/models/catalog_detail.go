package models

import (
	"time"

	"gorm.io/gorm"
)

type CatalogDetail struct {
	gorm.Model
	ID          uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CatalogID   uint       `json:"catalog_id" gorm:"column:catalog_id"`
	ItemID      uint       `json:"item_id" gorm:"column:item_id"`
	PriceBuy    int        `json:"price_buy" gorm:"column:price_buy"`
	PriceSell   int        `json:"price_sell" gorm:"column:price_sell"`
	Remark      string     `json:"remark" gorm:"column:remark"`
	IsActive    int8       `json:"is_active" gorm:"column:is_active"`
	CreatedByID *uint      `json:"created_by_id" gorm:"column:created_by_id"`
	UpdatedByID *uint      `json:"updated_by_id" gorm:"column:updated_by_id"`
	DeletedByID *uint      `json:"deleted_by_id" gorm:"column:deleted_by_id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
