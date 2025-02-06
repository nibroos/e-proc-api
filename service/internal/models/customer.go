package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID             uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CustomerTypeID uint       `json:"customer_type_id" gorm:"column:customer_type_id"`
	Name           string     `json:"name" gorm:"column:name"`
	Email          string     `json:"email" gorm:"column:email"`
	Phone          string     `json:"phone" gorm:"column:phone"`
	Address        string     `json:"address" gorm:"column:address"`
	Pic            string     `json:"pic" gorm:"column:pic"`
	IsActive       int8       `json:"is_active" gorm:"column:is_active"`
	UserID         uint       `json:"user_id" gorm:"column:user_id"`
	CreatedByID    *uint      `json:"created_by_id" gorm:"column:created_by_id"`
	UpdatedByID    *uint      `json:"updated_by_id" gorm:"column:updated_by_id"`
	DeletedByID    *uint      `json:"deleted_by_id" gorm:"column:deleted_by_id"`
	CreatedAt      *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt      *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
