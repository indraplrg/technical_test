package model

import (
	"time"
)

// Jurusan represents an academic department. One Jurusan has many Mahasiswa.
type Jurusan struct {
	ID          uint   `gorm:"column:id_jurusan;primaryKey" json:"id_jurusan"`
	NamaJurusan string `gorm:"size:100;not null" json:"nama_jurusan"`
	Fakultas    string `gorm:"size:100;not null" json:"fakultas"`
	Jenjang     string `gorm:"size:50;not null" json:"jenjang"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}
