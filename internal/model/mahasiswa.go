package model

import (
	"time"
)

// Mahasiswa represents a student. Every Mahasiswa belongs to one Jurusan.
type Mahasiswa struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Nama         string `gorm:"size:100;not null" json:"nama"`
	Umur         int    `gorm:"not null" json:"umur"`
	NIM          string `gorm:"size:20;not null;uniqueIndex" json:"nim"`
	TanggalLahir string `gorm:"not null" json:"tanggal_lahir"`
	Alamat       string `gorm:"type:text;not null" json:"alamat"`
	JurusanID    uint   `gorm:"column:id_jurusan;not null;index" json:"id_jurusan"`

	Jurusan *Jurusan `gorm:"foreignKey:JurusanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"jurusan,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
