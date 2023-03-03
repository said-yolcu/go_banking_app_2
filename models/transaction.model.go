package models

import (
	"time"
)

type Transaction struct {
	// Id represents the is of the transaction
	ID uint `gorm:"primaryKey;serial" json:"id"`
	// Base account
	UserId string `gorm:"varchar(11); uniqueIndex; not null" json:"user_id"`
	// Other account
	OtherId string `gorm:"varchar(11); uniqueIndex; not null" json:"other_id"`
	// Amount in dollars, from user account to other account
	Amount int `gorm:"not null" json:"amount"`
	// Each transaction belongs to two users
	// User User `gorm:"foreignKey:UserId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionReq struct {
	// Base account
	UserId string `gorm:"varchar(11); uniqueIndex; not null" json:"user_id"`
	// Other account
	OtherId string `gorm:"varchar(11); uniqueIndex; not null" json:"other_id"`
	// Amount in dollars, from user account to other account
	Amount int `gorm:"not null" json:"amount"`
}
