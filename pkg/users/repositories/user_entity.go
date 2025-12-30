// Package repositories provides specific way to store petshop data into databases.
package repositories

import "time"

type UserEntity struct {
	ID           string     `gorm:"primary_key;"`
	Username     string     `gorm:"size:100;not null;"`
	FirstName    string     `gorm:"size:100;not null;"`
	LastName     string     `gorm:"size:100;not null;"`
	Email        string     `gorm:"size:100;not null;"`
	PasswordHash string     `gorm:"size:255;not null;"`
	Status       string     `gorm:"size:100;not null;"`
	Role         string     `gorm:"size:100;not null;"`
	CreatedAt    time.Time  `gorm:"not null;"`
	UpdatedAt    time.Time  `gorm:"not null;"`
	DeletedAt    *time.Time ``
}

func (ue *UserEntity) TableName() string {
	return "users"
}
