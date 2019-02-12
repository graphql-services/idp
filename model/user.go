package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	// gorm.Model
	IID            uuid.UUID `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
	IEmail         string     `gorm:"unique"`
	IPasswordHash  string
	IEmailVerified bool
	IGivenName     *string
	IFamilyName    *string
	IMiddleName    *string
}

// UserInput ...
type UserInput struct {
	Email      string
	Password   string
	GivenName  *string
	FamilyName *string
	MiddleName *string
}
