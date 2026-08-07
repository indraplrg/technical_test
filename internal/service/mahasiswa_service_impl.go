package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/model"
	"github.com/indraplrg/technical_test/internal/repository"
)

const (
	defaultPage    = 1
	defaultLimit   = 10
	maxLimit       = 100
	dateOnlyLayout = "2006-01-02"
)

type mahasiswaService struct {
	mahasiswaRepo repository.MahasiswaRepository
	jurusanRepo   repository.JurusanRepository
}

// NewMahasiswaService builds the MahasiswaService with its dependencies.
func NewMahasiswaService(mahasiswaRepo repository.MahasiswaRepository, jurusanRepo repository.JurusanRepository) MahasiswaService {
	return &mahasiswaService{
		mahasiswaRepo: mahasiswaRepo,
		jurusanRepo:   jurusanRepo,
	}
}

func (s *mahasiswaService) Create(ctx context.Context, req dto.MahasiswaRequest) (*model.Mahasiswa, error) {
	if err := validateDateFormat(req.TanggalLahir); err != nil {
		return nil, err
	}
	if err := s.ensureJurusanExists(ctx, req.IDJurusan); err != nil {
		return nil, err
	}
	if err := s.ensureNIMUnique(ctx, req.NIM, 0); err != nil {
		return nil, err
	}

	mahasiswa := &model.Mahasiswa{
		Nama:         strings.TrimSpace(req.Nama),
		Umur:         req.Umur,
		NIM:          strings.TrimSpace(req.NIM),
		TanggalLahir: req.TanggalLahir,
		Alamat:       strings.TrimSpace(req.Alamat),
		IDJurusan:    req.IDJurusan,
	}
	if err := s.mahasiswaRepo.Create(ctx, mahasiswa); err != nil {
		return nil, err
	}
	return s.mahasiswaRepo.FindByID(ctx, mahasiswa.ID)
}

func (s *mahasiswaService) GetAll(ctx context.Context, query MahasiswaQuery) ([]model.Mahasiswa, *repository.Pagination, error) {
	if query.Page < 1 {
		query.Page = defaultPage
	}
	if query.Limit < 1 {
		query.Limit = defaultLimit
	}
	if query.Limit > maxLimit {
		query.Limit = maxLimit
	}

	filter := repository.MahasiswaFilter{
		Search:    query.Search,
		NIM:       query.NIM,
		IDJurusan: query.IDJurusan,
		SortBy:    query.SortBy,
		SortOrder: query.SortOrder,
	}
	return s.mahasiswaRepo.FindAll(ctx, filter, query.Page, query.Limit)
}

func (s *mahasiswaService) GetByID(ctx context.Context, id uint) (*model.Mahasiswa, error) {
	mahasiswa, err := s.mahasiswaRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, ErrMahasiswaNotFound
		}
		return nil, err
	}
	return mahasiswa, nil
}

func (s *mahasiswaService) Update(ctx context.Context, id uint, req dto.MahasiswaRequest) (*model.Mahasiswa, error) {
	if err := validateDateFormat(req.TanggalLahir); err != nil {
		return nil, err
	}

	mahasiswa, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := s.ensureJurusanExists(ctx, req.IDJurusan); err != nil {
		return nil, err
	}
	if err := s.ensureNIMUnique(ctx, req.NIM, id); err != nil {
		return nil, err
	}

	mahasiswa.Nama = strings.TrimSpace(req.Nama)
	mahasiswa.Umur = req.Umur
	mahasiswa.NIM = strings.TrimSpace(req.NIM)
	mahasiswa.TanggalLahir = req.TanggalLahir
	mahasiswa.Alamat = strings.TrimSpace(req.Alamat)
	mahasiswa.IDJurusan = req.IDJurusan

	if err := s.mahasiswaRepo.Update(ctx, mahasiswa); err != nil {
		return nil, err
	}
	return s.mahasiswaRepo.FindByID(ctx, id)
}

func (s *mahasiswaService) Delete(ctx context.Context, id uint) error {
	if _, err := s.GetByID(ctx, id); err != nil {
		return err
	}
	return s.mahasiswaRepo.Delete(ctx, id)
}

func (s *mahasiswaService) ensureJurusanExists(ctx context.Context, id uint) error {
	exists, err := s.jurusanRepo.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrJurusanNotFound
	}
	return nil
}

func (s *mahasiswaService) ensureNIMUnique(ctx context.Context, nim string, excludeID uint) error {
	existing, err := s.mahasiswaRepo.FindByNIM(ctx, strings.TrimSpace(nim), excludeID)
	if err != nil && !errors.Is(err, repository.ErrRecordNotFound) {
		return err
	}
	if existing != nil && existing.ID != excludeID {
		return ErrDuplicateNIM
	}
	return nil
}

func validateDateFormat(value string) error {
	if _, err := time.Parse(dateOnlyLayout, strings.TrimSpace(value)); err != nil {
		return ErrInvalidDate
	}
	return nil
}
