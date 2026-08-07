package database

import (
	"log/slog"

	"github.com/indraplrg/technical_test/internal/model"
	"gorm.io/gorm"
)

// Seed inserts starter data when tables are empty.
func Seed(db *gorm.DB) error {
	var jurusanCount int64
	var mahasiswaCount int64

	if err := db.Model(&model.Jurusan{}).Count(&jurusanCount).Error; err != nil {
		return err
	}

	if err := db.Model(&model.Mahasiswa{}).Count(&mahasiswaCount).Error; err != nil {
		return err
	}
	
	if jurusanCount > 0 || mahasiswaCount > 0 {
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

		jurusanMap := make(map[string]uint)
		for _, j := range jurusanList {
			jurusanMap[j.NamaJurusan] = j.ID
		}

		mahasiswaList := []model.Mahasiswa{
		{
			Nama:         "Andi Pratama",
			Umur:         20,
			NIM:          "2201001",
			TanggalLahir: "2004-01-15",
			Alamat:       "Jl. Merdeka No. 10, Jakarta",
			JurusanID:    jurusanMap["Teknik Informatika"],
		},
		{
			Nama:         "Budi Santoso",
			Umur:         21,
			NIM:          "2201002",
			TanggalLahir: "2003-05-22",
			Alamat:       "Jl. Sudirman No. 25, Bandung",
			JurusanID:    jurusanMap["Teknik Informatika"],
		},
		{
			Nama:         "Citra Lestari",
			Umur:         20,
			NIM:          "2202001",
			TanggalLahir: "2004-03-10",
			Alamat:       "Jl. Diponegoro No. 8, Surabaya",
			JurusanID:    jurusanMap["Sistem Informasi"],
		},
		{
			Nama:         "Dewi Anggraini",
			Umur:         22,
			NIM:          "2202002",
			TanggalLahir: "2002-11-02",
			Alamat:       "Jl. Ahmad Yani No. 17, Semarang",
			JurusanID:    jurusanMap["Sistem Informasi"],
		},
		{
			Nama:         "Eko Saputra",
			Umur:         21,
			NIM:          "2203001",
			TanggalLahir: "2003-07-18",
			Alamat:       "Jl. Gatot Subroto No. 3, Medan",
			JurusanID:    jurusanMap["Teknik Elektro"],
		},
		{
			Nama:         "Fajar Nugroho",
			Umur:         20,
			NIM:          "2203002",
			TanggalLahir: "2004-09-05",
			Alamat:       "Jl. Pahlawan No. 11, Yogyakarta",
			JurusanID:    jurusanMap["Teknik Elektro"],
		},
		{
			Nama:         "Gita Permata",
			Umur:         19,
			NIM:          "2204001",
			TanggalLahir: "2005-02-28",
			Alamat:       "Jl. Kartini No. 6, Malang",
			JurusanID:    jurusanMap["Akuntansi"],
		},
		{
			Nama:         "Hendra Wijaya",
			Umur:         21,
			NIM:          "2204002",
			TanggalLahir: "2003-10-14",
			Alamat:       "Jl. Veteran No. 20, Palembang",
			JurusanID:    jurusanMap["Akuntansi"],
		},
		{
			Nama:         "Intan Maharani",
			Umur:         20,
			NIM:          "2205001",
			TanggalLahir: "2004-06-08",
			Alamat:       "Jl. Imam Bonjol No. 9, Makassar",
			JurusanID:    jurusanMap["Manajemen"],
		},
		{
			Nama:         "Joko Firmansyah",
			Umur:         22,
			NIM:          "2205002",
			TanggalLahir: "2002-12-21",
			Alamat:       "Jl. Sisingamangaraja No. 15, Denpasar",
			JurusanID:    jurusanMap["Manajemen"],
		},
	}

	if err := db.Create(&mahasiswaList).Error; err != nil {
		return err
	}

	slog.Info("seeded jurusan data", "count", len(jurusanList))
	slog.Info("seeded mahasiswa data", "count", len(mahasiswaList))
	return nil
}
