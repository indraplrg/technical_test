package controller

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/service"
)

// ExportController exposes file export endpoints for Mahasiswa.
type ExportController struct {
	service service.MahasiswaService
}

// NewExportController builds the export controller with its dependency.
func NewExportController(service service.MahasiswaService) *ExportController {
	return &ExportController{service: service}
}

// ExportCSV handles GET /api/v1/mahasiswa/export/csv.
func (ctr *ExportController) ExportCSV(c *gin.Context) {
	list, err := ctr.service.ExportAll(c.Request.Context(), ctr.buildQuery(c))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	filename := "mahasiswa_" + time.Now().UTC().Format("20060102150405") + ".csv"
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Writer.WriteHeader(http.StatusOK)

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	_ = writer.Write([]string{"ID", "Nama", "Umur", "NIM", "TanggalLahir", "Alamat", "IDJurusan", "Jurusan"})
	for _, m := range list {
		jurusan := ""
		if m.Jurusan != nil {
			jurusan = m.Jurusan.NamaJurusan
		}
		_ = writer.Write([]string{
			strconv.FormatUint(uint64(m.ID), 10),
			m.Nama,
			strconv.Itoa(m.Umur),
			m.NIM,
			m.TanggalLahir,
			m.Alamat,
			strconv.FormatUint(uint64(m.IDJurusan), 10),
			jurusan,
		})
	}
}

// buildQuery reads export filter query parameters.
func (ctr *ExportController) buildQuery(c *gin.Context) service.MahasiswaQuery {
	query := service.MahasiswaQuery{
		Search: c.Query("search"),
		NIM:    c.Query("nim"),
	}
	if idJurusan, err := parseUint(c.Query("id_jurusan")); err == nil {
		query.IDJurusan = idJurusan
	}
	return query
}
