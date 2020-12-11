package entity

import "time"

type Pet struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	Status    string     `gorm:"size:100;not null;" json:"status"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Pets []Pet

func (p Pet) Prepare() {
	p.UpdatedAt = time.Now()
	p.CreatedAt = time.Now()
}
