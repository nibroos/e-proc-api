package models

import (
	"time"

	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model
	ID          uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CustomerID  uint       `json:"customer_id" gorm:"column:customer_id"`
	CatalogNo   string     `json:"catalog_no" gorm:"column:catalog_no"`
	Description string     `json:"description" gorm:"column:description"`
	Remark      string     `json:"remark" gorm:"column:remark"`
	IsActive    int8       `json:"is_active" gorm:"column:is_active"`
	CreatedByID *uint      `json:"created_by_id" gorm:"column:created_by_id"`
	UpdatedByID *uint      `json:"updated_by_id" gorm:"column:updated_by_id"`
	DeletedByID *uint      `json:"deleted_by_id" gorm:"column:deleted_by_id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
