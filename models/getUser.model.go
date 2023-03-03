package models

type GetUser struct {
	Name    string `gorm:"varchar(255); not null" json:"name"`
	Surname string `gorm:"varchar(255); not null" json:"surname"`
}
