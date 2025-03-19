package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Role      string    `gorm:"type:varchar(50);not null;check:role IN ('admin', 'manager', 'staff')" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:null" json:"updated_at"`
}
