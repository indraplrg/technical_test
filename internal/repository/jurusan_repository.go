package repository

import (
	"context"
	"errors"

	"github.com/indraplrg/technical_test/internal/model"
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

// JurusanRepository defines data access operations for Jurusan.
type JurusanRepository interface {
	Create(ctx context.Context, jurusan *model.Jurusan) error
	FindAll(ctx context.Context) ([]model.Jurusan, error)
	FindByID(ctx context.Context, id uint) (*model.Jurusan, error)
	Update(ctx context.Context, jurusan *model.Jurusan) error
	Delete(ctx context.Context, id uint) error
	ExistsByID(ctx context.Context, id uint) (bool, error)
	ExistsByName(ctx context.Context, name string, excludeID uint) (bool, error)
}

var ErrDeleteHasChildren = errors.New("record has related mahasiswa records")