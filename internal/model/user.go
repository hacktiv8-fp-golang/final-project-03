package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name" valid:"required"`
	Email     string    `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password  string    `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
	Role      string    `gorm:"not null" json:"role" valid:"required,in(admin|member)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Task
}

type UserCreate struct {
	FullName string `gorm:"not null" json:"full_name" valid:"required"`
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}

type UserUpdate struct {
	FullName string `gorm:"not null" json:"full_name" valid:"required"`
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
}

type LoginCredential struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}
