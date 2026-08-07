package service

import "errors"

// Domain errors returned by the service layer.
// Controllers map these to proper HTTP status codes.
var (
	ErrNotFound             = errors.New("resource not found")
	ErrDuplicateNamaJurusan = errors.New("nama jurusan already exists")
	ErrDuplicateNIM         = errors.New("nim already exists")
	ErrJurusanNotFound      = errors.New("jurusan not found")
	ErrJurusanHasMahasiswa  = errors.New("jurusan still has associated mahasiswa")
	ErrMahasiswaNotFound    = errors.New("mahasiswa not found")
	ErrInvalidDate          = errors.New("tanggal_lahir must use YYYY-MM-DD format")
)
