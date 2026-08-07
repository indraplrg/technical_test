package service

import (
	"context"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/model"
)

// JurusanService contains business logic for Jurusan.
type JurusanService interface {
	Create(ctx context.Context, req dto.JurusanRequest) (*model.Jurusan, error)
	GetAll(ctx context.Context) ([]model.Jurusan, error)
	GetByID(ctx context.Context, id uint) (*model.Jurusan, error)
	Update(ctx context.Context, id uint, req dto.JurusanRequest) (*model.Jurusan, error)
	Delete(ctx context.Context, id uint) error
}
