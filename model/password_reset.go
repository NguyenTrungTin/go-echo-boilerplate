package model

import "time"

type PasswordReset struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `gorm:"" json:"user_id"`
	Email     string     `gorm:"not null" json:"email"`
	Token     string     `gorm:"" json:"token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
