package database

import (
	"log/slog"

	"github.com/indraplrg/technical_test/internal/model"
	"gorm.io/gorm"
)

// Seed inserts starter data when tables are empty.
func Seed(db *gorm.DB) error {
	var jurusanCount int64
	if err := db.Model(&model.Jurusan{}).Count(&jurusanCount).Error; err != nil {
		return err
	}
	if jurusanCount > 0 {
		return nil
	}

	jurusanList := []model.Jurusan{
		{NamaJurusan: "Teknik Informatika", Fakultas: "Fakultas Ilmu Komputer", Jenjang: "S1"},
		{NamaJurusan: "Sistem Informasi", Fakultas: "Fakultas Ilmu Komputer", Jenjang: "S1"},
		{NamaJurusan: "Teknik Elektro", Fakultas: "Fakultas Teknik", Jenjang: "S1"},
		{NamaJurusan: "Akuntansi", Fakultas: "Fakultas Ekonomi", Jenjang: "D3"},
		{NamaJurusan: "Manajemen", Fakultas: "Fakultas Ekonomi", Jenjang: "S1"},
	}

	if err := db.Create(&jurusanList).Error; err != nil {
		return err
	}
	slog.Info("seeded jurusan data", "count", len(jurusanList))
	return nil
}
