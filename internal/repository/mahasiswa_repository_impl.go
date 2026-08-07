package repository

import (
	"context"
	"strings"

	"github.com/indraplrg/technical_test/internal/model"
	"gorm.io/gorm"
)

// sortableColumns whitelists allowed sort fields to prevent SQL injection.
var sortableColumns = map[string]string{
	"nama":          "nama",
	"umur":          "umur",
	"nim":           "nim",
	"tanggal_lahir": "tanggal_lahir",
	"created_at":    "created_at",
}

type mahasiswaRepository struct {
	db *gorm.DB
}

// NewMahasiswaRepository creates a MahasiswaRepository backed by the given DB.
func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db: db}
}

func (r *mahasiswaRepository) Create(ctx context.Context, mahasiswa *model.Mahasiswa) error {
	return r.db.WithContext(ctx).Create(mahasiswa).Error
}

func (r *mahasiswaRepository) FindAll(ctx context.Context, filter MahasiswaFilter, page, limit int) ([]model.Mahasiswa, *Pagination, error) {
	query := r.applyFilter(ctx, filter)

	var total int64
	if err := query.Model(&model.Mahasiswa{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	query = r.applySort(query, filter)

	var mahasiswaList []model.Mahasiswa
	err := query.
		Preload("Jurusan").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&mahasiswaList).Error
	if err != nil {
		return nil, nil, err
	}

	pagination := &Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: int((total + int64(limit) - 1) / int64(limit)),
	}
	return mahasiswaList, pagination, nil
}

func (r *mahasiswaRepository) FindByID(ctx context.Context, id uint) (*model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	err := r.db.WithContext(ctx).
		Preload("Jurusan").
		First(&mahasiswa, id).Error
	return &mahasiswa, err
}

func (r *mahasiswaRepository) Update(ctx context.Context, mahasiswa *model.Mahasiswa) error {
	return r.db.WithContext(ctx).
		Model(mahasiswa).
		Select("nama", "umur", "nim", "tanggal_lahir", "alamat", "id_jurusan").
		Updates(mahasiswa).Error
}

func (r *mahasiswaRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Mahasiswa{}, id).Error
}

func (r *mahasiswaRepository) FindAllExport(ctx context.Context, filter MahasiswaFilter) ([]model.Mahasiswa, error) {
	var mahasiswaList []model.Mahasiswa
	err := r.applyFilter(ctx, filter).
		Preload("Jurusan").
		Order("nama ASC").
		Find(&mahasiswaList).Error
	return mahasiswaList, err
}

func (r *mahasiswaRepository) FindByNIM(ctx context.Context, nim string, excludeID uint) (*model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	query := r.db.WithContext(ctx).Where("nim = ?", nim)
	if excludeID != 0 {
		query = query.Where("id <> ?", excludeID)
	}
	err := query.First(&mahasiswa).Error
	return &mahasiswa, err
}

func (r *mahasiswaRepository) ExistsByID(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Mahasiswa{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *mahasiswaRepository) CountByJurusan(ctx context.Context, idJurusan uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Mahasiswa{}).
		Where("id_jurusan = ?", idJurusan).
		Count(&count).Error
	return count, err
}

func (r *mahasiswaRepository) applyFilter(ctx context.Context, filter MahasiswaFilter) *gorm.DB {
	query := r.db.WithContext(ctx).Model(&model.Mahasiswa{})

	if search := strings.TrimSpace(filter.Search); search != "" {
		pattern := "%" + search + "%"
		query = query.Where("nama ILIKE ? OR nim ILIKE ?", pattern, pattern)
	}
	if nim := strings.TrimSpace(filter.NIM); nim != "" {
		query = query.Where("nim = ?", nim)
	}
	if filter.IDJurusan != 0 {
		query = query.Where("id_jurusan = ?", filter.IDJurusan)
	}
	return query
}

func (r *mahasiswaRepository) applySort(query *gorm.DB, filter MahasiswaFilter) *gorm.DB {
	column, ok := sortableColumns[strings.ToLower(strings.TrimSpace(filter.SortBy))]
	if !ok {
		column = "created_at"
	}
	order := strings.ToUpper(strings.TrimSpace(filter.SortOrder))
	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}
	return query.Order(column + " " + order)
}
