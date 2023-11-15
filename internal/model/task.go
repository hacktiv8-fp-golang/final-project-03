package model

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey" `
	Title       string    `json:"title" gorm:"not null" valid:"required"`
	Description string    `json:"description" gorm:"not null" valid:"required"`
	Status      bool      `json:"status" gorm:"not null" valid:"required,in(admin|member)"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        *User
}

type TaskCreate struct {
	Title       string `json:"title" gorm:"not null" valid:"required"`
	Description string `json:"description" gorm:"not null" valid:"required"`
	CategoryID  uint   `json:"category_id"`
}

type TaskUpdate struct {
	Title       string `json:"title" gorm:"not null" valid:"required"`
	Description string `json:"description" gorm:"not null" valid:"required"`
	CategoryID  uint   `json:"category_id"`
}

type TaskStatusUpdate struct {
	Status bool `json:"status" gorm:"not null" valid:"required,in(admin|member)"`
}

type TaskCategoryUpdate struct {
	CategoryID uint `json:"category_id"`
}
