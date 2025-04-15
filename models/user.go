package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"unique" json:"username"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"password"`          // jangan tampilkan password
	Obats     []Obat     `gorm:"foreignKey:UserID"` // relasi ke obat
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
