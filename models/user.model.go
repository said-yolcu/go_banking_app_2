package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("Binfeqr_gu5wn_bhiHdmhr_gvcOub8nn_biLbr_cofcruk")

// Denotes the user
type User struct {
	// Created a universally unique identifier for ID
	// UserID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key; uniqueIndex" json:"user_id"`

	Name      string `gorm:"varchar(255); not null" json:"name"`
	Surname   string `gorm:"varchar(255); not null" json:"surname"`
	StateID   string `gorm:"varchar(11); uniqueIndex; primary_key; not null" json:"state_id"`
	Email     string `gorm:"varchar(255); uniqueIndex; not null" json:"email"`
	Phone     string `gorm:"varchar(255); uniqueIndex; not null" json:"phone"`
	Password  string `gorm:"varchar(255); not null" json:"password"`
	Balance   int    `gorm:"not null" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// It is used at login
type Credentials struct {
	StateID  string `json:"state_id"`
	Password string `json:"password"`
}

type Claims struct {
	StateID string
	jwt.StandardClaims
}
