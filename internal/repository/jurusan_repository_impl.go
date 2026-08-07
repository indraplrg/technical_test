package repository

import (
	"context"

	"github.com/indraplrg/technical_test/internal/model"
	"gorm.io/gorm"
)

type jurusanRepository struct {
	db *gorm.DB
}

// NewJurusanRepository creates a JurusanRepository backed by the given DB.
func NewJurusanRepository(db *gorm.DB) JurusanRepository {
	return &jurusanRepository{db: db}
}

func (r *jurusanRepository) Create(ctx context.Context, jurusan *model.Jurusan) error {
	return r.db.WithContext(ctx).Create(jurusan).Error
}

func (r *jurusanRepository) FindAll(ctx context.Context) ([]model.Jurusan, error) {
	var jurusanList []model.Jurusan
	err := r.db.WithContext(ctx).
		Order("nama_jurusan ASC").
		Find(&jurusanList).Error
	return jurusanList, err
}

func (r *jurusanRepository) FindByID(ctx context.Context, id uint) (*model.Jurusan, error) {
	var jurusan model.Jurusan
	err := r.db.WithContext(ctx).First(&jurusan, id).Error
	return &jurusan, err
}

func (r *jurusanRepository) Update(ctx context.Context, jurusan *model.Jurusan) error {
	return r.db.WithContext(ctx).
		Model(jurusan).
		Select("nama_jurusan", "fakultas", "jenjang").
		Updates(jurusan).Error
}

func (r *jurusanRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Jurusan{}, id).Error
}

func (r *jurusanRepository) ExistsByID(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Jurusan{}).
		Where("id_jurusan = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *jurusanRepository) ExistsByName(ctx context.Context, name string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).
		Model(&model.Jurusan{}).
		Where("nama_jurusan = ?", name)
	if excludeID != 0 {
		query = query.Where("id_jurusan <> ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}