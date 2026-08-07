package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/response"
	"github.com/indraplrg/technical_test/internal/service"
	"github.com/indraplrg/technical_test/internal/validator"
)

// JurusanController handles Jurusan HTTP requests. It stays thin:
// all business logic lives in the service layer.
type JurusanController struct {
	service service.JurusanService
}

// NewJurusanController builds the controller with its service dependency.
func NewJurusanController(service service.JurusanService) *JurusanController {
	return &JurusanController{service: service}
}

// Create handles POST /api/v1/jurusan.
func (ctr *JurusanController) Create(c *gin.Context) {
	var req dto.JurusanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", validator.ValidateErrors(err))
		return
	}

	jurusan, err := ctr.service.Create(c.Request.Context(), req)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, "jurusan created", jurusan)
}

// GetAll handles GET /api/v1/jurusan.
func (ctr *JurusanController) GetAll(c *gin.Context) {
	jurusanList, err := ctr.service.GetAll(c.Request.Context())
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "jurusan fetched successfully", jurusanList)
}

// GetByID handles GET /api/v1/jurusan/:id.
func (ctr *JurusanController) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	jurusan, err := ctr.service.GetByID(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "jurusan fetched successfully", jurusan)
}

// Update handles PUT /api/v1/jurusan/:id.
func (ctr *JurusanController) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var req dto.JurusanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", validator.ValidateErrors(err))
		return
	}

	jurusan, err := ctr.service.Update(c.Request.Context(), id, req)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "jurusan updated successfully", jurusan)
}

// Delete handles DELETE /api/v1/jurusan/:id.
func (ctr *JurusanController) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := ctr.service.Delete(c.Request.Context(), id); err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "jurusan deleted successfully", nil)
}
