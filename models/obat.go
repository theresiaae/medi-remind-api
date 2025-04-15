package models

import "time"

type Obat struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	NamaObat   string     `json:"nama" gorm:"column:nama_obat"`
	Dosis      string     `json:"dosis" gorm:"column:dosis"`
	WaktuMinum string     `json:"waktu" gorm:"column:waktu_minum"`
	Keterangan string     `json:"deskripsi" gorm:"column:keterangan"`
	UserID     uint       `json:"user_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
