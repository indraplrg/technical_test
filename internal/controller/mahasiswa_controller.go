package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/dto"
	"github.com/indraplrg/technical_test/internal/response"
	"github.com/indraplrg/technical_test/internal/service"
	"github.com/indraplrg/technical_test/internal/validator"
)

// MahasiswaController handles Mahasiswa HTTP endpoints. It stays thin:
// all business logic lives in the service layer.
type MahasiswaController struct {
	service service.MahasiswaService
}

// NewMahasiswaController builds the controller with its service dependency.
func NewMahasiswaController(service service.MahasiswaService) *MahasiswaController {
	return &MahasiswaController{service: service}
}

// Create handles POST /api/v1/mahasiswa.
// @Summary Create a new mahasiswa
// @Description Create a new mahasiswa record
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param body body dto.MahasiswaRequest true "Mahasiswa payload"
// @Success 201 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Failure 409 {object} response.Result
// @Router /mahasiswa [post]
func (ctr *MahasiswaController) Create(c *gin.Context) {
	var req dto.MahasiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", validator.ValidateErrors(err))
		return
	}

	mahasiswa, err := ctr.service.Create(c.Request.Context(), req)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, "mahasiswa created", mahasiswa)
}

// GetAll handles GET /api/v1/mahasiswa with search, filter, sort
// and pagination support.
// @Summary Get all mahasiswa
// @Description Returns a paginated list of mahasiswa with search, filter and sort
// @Tags mahasiswa
// @Produce json
// @Param search query string false "Search by nama or nim"
// @Param nim query string false "Filter by exact nim"
// @Param id_jurusan query int false "Filter by jurusan id"
// @Param sort_by query string false "Sort field: nama, umur, nim, tanggal_lahir, created_at"
// @Param sort_order query string false "Sort direction: asc or desc"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} response.Result
// @Failure 500 {object} response.Result
// @Router /mahasiswa [get]
func (ctr *MahasiswaController) GetAll(c *gin.Context) {
	query := service.MahasiswaQuery{
		Search:    c.Query("search"),
		NIM:       c.Query("nim"),
		SortBy:    c.Query("sort_by"),
		SortOrder: c.Query("sort_order"),
	}

	if idJurusan, err := parseUint(c.Query("id_jurusan")); err == nil {
		query.IDJurusan = idJurusan
	}
	query.Page = parseOrDefault(c.Query("page"), 1)
	query.Limit = parseOrDefault(c.Query("limit"), 10)

	mahasiswaList, pagination, err := ctr.service.GetAll(c.Request.Context(), query)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	data := response.Paginated{
		Items:      mahasiswaList,
		Pagination: pagination,
	}
	response.Success(c, http.StatusOK, "mahasiswa fetched successfully", data)
}

// GetByID handles GET /api/v1/mahasiswa/:id.
// @Summary Get mahasiswa by id
// @Description Returns a single mahasiswa record
// @Tags mahasiswa
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Success 200 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Router /mahasiswa/{id} [get]
func (ctr *MahasiswaController) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	mahasiswa, err := ctr.service.GetByID(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "mahasiswa fetched successfully", mahasiswa)
}

// Update handles PUT /api/v1/mahasiswa/:id.
// @Summary Update mahasiswa
// @Description Update an existing mahasiswa record
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Param body body dto.MahasiswaRequest true "Mahasiswa payload"
// @Success 200 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Failure 409 {object} response.Result
// @Router /mahasiswa/{id} [put]
func (ctr *MahasiswaController) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var req dto.MahasiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", validator.ValidateErrors(err))
		return
	}

	mahasiswa, err := ctr.service.Update(c.Request.Context(), id, req)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "mahasiswa updated successfully", mahasiswa)
}

// Delete handles DELETE /api/v1/mahasiswa/:id.
// @Summary Delete mahasiswa
// @Description Delete a mahasiswa record (soft delete)
// @Tags mahasiswa
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Success 200 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Router /mahasiswa/{id} [delete]
func (ctr *MahasiswaController) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := ctr.service.Delete(c.Request.Context(), id); err != nil {
		handleServiceError(c, err)
		return
	}
	response.Success(c, http.StatusOK, "mahasiswa deleted successfully", nil)
}
