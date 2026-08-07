package service

import (
	"context"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/model"
	"github.com/indraplrg/technical_test/internal/repository"
)

// MahasiswaQuery carries list/search/filter/sort/pagination parameters.
type MahasiswaQuery struct {
	Search    string
	NIM       string
	IDJurusan uint
	SortBy    string
	SortOrder string
	Page      int
	Limit     int
}

// MahasiswaService contains business logic for Mahasiswa.
type MahasiswaService interface {
	Create(ctx context.Context, req dto.MahasiswaRequest) (*model.Mahasiswa, error)
	GetAll(ctx context.Context, query MahasiswaQuery) ([]model.Mahasiswa, *repository.Pagination, error)
	GetByID(ctx context.Context, id uint) (*model.Mahasiswa, error)
	Update(ctx context.Context, id uint, req dto.MahasiswaRequest) (*model.Mahasiswa, error)
	Delete(ctx context.Context, id uint) error
	ExportAll(ctx context.Context, query MahasiswaQuery) ([]model.Mahasiswa, error)
}
