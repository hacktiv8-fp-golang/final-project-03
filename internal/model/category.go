package model

import (
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Type      string    `json:"type" gorm:"not null" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task
}

type CategoryCreate struct {
	Type string `json:"type" gorm:"not null" valid:"required"`
}

type CategoryUpdate struct {
	Type string `json:"type" gorm:"not null" valid:"required"`
}
