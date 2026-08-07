package repository

import (
	"context"

	"github.com/indraplrg/technical_test/internal/model"
)

// Pagination holds pagination metadata returned by list operations.
type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_pages"`
}

// MahasiswaFilter describes search and filter criteria.
type MahasiswaFilter struct {
	Search    string
	NIM       string
	IDJurusan uint
	SortBy    string
	SortOrder string
}

// MahasiswaRepository defines data access operations for Mahasiswa.
type MahasiswaRepository interface {
	Create(ctx context.Context, mahasiswa *model.Mahasiswa) error
	FindAll(ctx context.Context, filter MahasiswaFilter, page, limit int) ([]model.Mahasiswa, *Pagination, error)
	FindByID(ctx context.Context, id uint) (*model.Mahasiswa, error)
	Update(ctx context.Context, mahasiswa *model.Mahasiswa) error
	Delete(ctx context.Context, id uint) error
	FindAllExport(ctx context.Context, filter MahasiswaFilter) ([]model.Mahasiswa, error)
	FindByNIM(ctx context.Context, nim string, excludeID uint) (*model.Mahasiswa, error)
	ExistsByID(ctx context.Context, id uint) (bool, error)
	CountByJurusan(ctx context.Context, idJurusan uint) (int64, error)
}
