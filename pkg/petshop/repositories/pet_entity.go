// Package repositories provides specific way to store petshop data into databases.
package repositories

import "time"

type PetEntity struct {
	ID        string     `gorm:"primary_key;"`
	Name      string     `gorm:"size:100;not null;"`
	Status    string     `gorm:"size:100;not null;"`
	CreatedAt time.Time  `gorm:"not null;"`
	UpdatedAt time.Time  `gorm:"not null;"`
	DeletedAt *time.Time ``
}

func (pe *PetEntity) TableName() string {
	return "pets"
}
