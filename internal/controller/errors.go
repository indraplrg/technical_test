package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/response"
	"github.com/indraplrg/technical_test/internal/service"
)

// handleServiceError maps service layer errors to proper HTTP responses.
func handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrNotFound):
		response.Error(c, http.StatusNotFound, err.Error())
	case errors.Is(err, service.ErrJurusanNotFound):
		response.Error(c, http.StatusNotFound, service.ErrJurusanNotFound.Error())
	case errors.Is(err, service.ErrMahasiswaNotFound):
		response.Error(c, http.StatusNotFound, service.ErrMahasiswaNotFound.Error())
	case errors.Is(err, service.ErrDuplicateNamaJurusan):
		response.Error(c, http.StatusConflict, service.ErrDuplicateNamaJurusan.Error())
	case errors.Is(err, service.ErrDuplicateNIM):
		response.Error(c, http.StatusConflict, service.ErrDuplicateNIM.Error())
	case errors.Is(err, service.ErrJurusanHasMahasiswa):
		response.Error(c, http.StatusConflict, service.ErrJurusanHasMahasiswa.Error())
	case errors.Is(err, service.ErrInvalidDate):
		response.Error(c, http.StatusBadRequest, service.ErrInvalidDate.Error())
	default:
		response.Error(c, http.StatusInternalServerError, "internal server error")
	}
}
