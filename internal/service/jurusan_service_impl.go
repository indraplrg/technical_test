package service

import (
	"context"
	"errors"
	"strings"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/model"
	"github.com/indraplrg/technical_test/internal/repository"
)

type jurusanService struct {
	jurusanRepo   repository.JurusanRepository
	mahasiswaRepo repository.MahasiswaRepository
}

// NewJurusanService builds the JurusanService with its dependencies.
func NewJurusanService(jurusanRepo repository.JurusanRepository, mahasiswaRepo repository.MahasiswaRepository) JurusanService {
	return &jurusanService{
		jurusanRepo:   jurusanRepo,
		mahasiswaRepo: mahasiswaRepo,
	}
}

func (s *jurusanService) Create(ctx context.Context, req dto.JurusanRequest) (*model.Jurusan, error) {
	name := strings.TrimSpace(req.NamaJurusan)
	exists, err := s.jurusanRepo.ExistsByName(ctx, name, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateNamaJurusan
	}

	jurusan := &model.Jurusan{
		NamaJurusan: name,
		Fakultas:    strings.TrimSpace(req.Fakultas),
		Jenjang:     strings.TrimSpace(req.Jenjang),
	}
	if err := s.jurusanRepo.Create(ctx, jurusan); err != nil {
		return nil, err
	}
	return jurusan, nil
}

func (s *jurusanService) GetAll(ctx context.Context) ([]model.Jurusan, error) {
	return s.jurusanRepo.FindAll(ctx)
}

func (s *jurusanService) GetByID(ctx context.Context, id uint) (*model.Jurusan, error) {
	jurusan, err := s.jurusanRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, ErrJurusanNotFound
		}
		return nil, err
	}
	return jurusan, nil
}

func (s *jurusanService) Update(ctx context.Context, id uint, req dto.JurusanRequest) (*model.Jurusan, error) {
	jurusan, err := s.jurusanRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, ErrJurusanNotFound
		}
		return nil, err
	}

	name := strings.TrimSpace(req.NamaJurusan)
	duplicate, err := s.jurusanRepo.ExistsByName(ctx, name, id)
	if err != nil {
		return nil, err
	}
	if duplicate {
		return nil, ErrDuplicateNamaJurusan
	}

	jurusan.NamaJurusan = name
	jurusan.Fakultas = strings.TrimSpace(req.Fakultas)
	jurusan.Jenjang = strings.TrimSpace(req.Jenjang)

	if err := s.jurusanRepo.Update(ctx, jurusan); err != nil {
		return nil, err
	}
	return jurusan, nil
}

func (s *jurusanService) Delete(ctx context.Context, id uint) error {
	if _, err := s.GetByID(ctx, id); err != nil {
		return err
	}

	count, err := s.mahasiswaRepo.CountByJurusan(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrJurusanHasMahasiswa
	}

	return s.jurusanRepo.Delete(ctx, id)
}
