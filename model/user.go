package model

import "time"

type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	FirstName    string     `gorm:"not null" json:"firstname" validate:"required"`
	LastName     string     `gorm:"not null" json:"lastname" validate:"required"`
	Username     string     `gorm:"not null;unique" json:"username" validate:"required"`
	Email        string     `gorm:"not null;unique_index" json:"email" validate:"required,email"`
	Login        string     `gorm:"-" json:"login,omitempty"`
	Role         string     `json:"role" csv:"ROLE,role"`
	Password     string     `gorm:"-" json:"password,omitempty" validate:"min=6"`
	PasswordHash string     `gorm:"not null" json:"-"`
	Status       string     `gorm:"default:'ACTIVE'" json:"status"`
	Timezone     string     `gorm:"default:NULL" json:"timezone"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
}
