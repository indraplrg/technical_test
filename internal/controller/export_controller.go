package controller

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pdf/fpdf"
	"github.com/xuri/excelize/v2"

	"github.com/indraplrg/technical_test/internal/response"
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

// ExportExcel handles GET /api/v1/mahasiswa/export/excel.
func (ctr *ExportController) ExportExcel(c *gin.Context) {
	list, err := ctr.service.ExportAll(c.Request.Context(), ctr.buildQuery(c))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	sheetName := "Mahasiswa"
	file := excelize.NewFile()
	defer file.Close()

	sheet := file.GetSheetName(file.GetActiveSheetIndex())
	file.SetSheetName(sheet, sheetName)

	headers := []interface{}{"ID", "Nama", "Umur", "NIM", "Tanggal Lahir", "Alamat", "ID Jurusan", "Jurusan"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = file.SetCellValue(sheetName, cell, header)
	}

	for idx, m := range list {
		row := idx + 2
		jurusan := ""
		if m.Jurusan != nil {
			jurusan = m.Jurusan.NamaJurusan
		}
		values := []interface{}{
			m.ID,
			m.Nama,
			m.Umur,
			m.NIM,
			m.TanggalLahir,
			m.Alamat,
			m.IDJurusan,
			jurusan,
		}
		for i, value := range values {
			cell, _ := excelize.CoordinatesToCellName(i+1, row)
			_ = file.SetCellValue(sheetName, cell, value)
		}
	}

	filename := "mahasiswa_" + time.Now().UTC().Format("20060102150405") + ".xlsx"
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Writer.WriteHeader(http.StatusOK)

	if _, err := file.WriteTo(c.Writer); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate excel file")
	}
}

// ExportPDF handles GET /api/v1/mahasiswa/export/pdf.
func (ctr *ExportController) ExportPDF(c *gin.Context) {
	list, err := ctr.service.ExportAll(c.Request.Context(), ctr.buildQuery(c))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Data Mahasiswa", "0", 1, "C", false, 0, "")
	pdf.Ln(4)

	headers := []string{"ID", "Nama", "Umur", "NIM", "Tgl Lahir", "Alamat", "Jurusan"}
	widths := []float64{14, 45, 12, 18, 24, 70, 40}

	pdf.SetFont("Arial", "B", 9)
	for i, header := range headers {
		pdf.CellFormat(widths[i], 8, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 9)
	for _, m := range list {
		jurusan := ""
		if m.Jurusan != nil {
			jurusan = m.Jurusan.NamaJurusan
		}
		row := []string{
			strconv.Itoa(int(m.ID)),
			m.Nama,
			strconv.Itoa(m.Umur),
			m.NIM,
			m.TanggalLahir,
			truncate(m.Alamat, 70),
			jurusan,
		}
		for i, cell := range row {
			pdf.CellFormat(widths[i], 8, cell, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}

	filename := "mahasiswa_" + time.Now().UTC().Format("20060102150405") + ".pdf"
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Writer.WriteHeader(http.StatusOK)

	if err := pdf.Output(c.Writer); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate pdf file")
	}
}

// ExportJSON handles GET /api/v1/mahasiswa/export/json.
func (ctr *ExportController) ExportJSON(c *gin.Context) {
	list, err := ctr.service.ExportAll(c.Request.Context(), ctr.buildQuery(c))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	payload := response.Result{
		Success: true,
		Message: "mahasiswa exported successfully",
		Data:    list,
	}

	raw, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to marshal json")
		return
	}

	filename := "mahasiswa_" + time.Now().UTC().Format("20060102150405") + ".json"
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Writer.WriteHeader(http.StatusOK)
	if _, err := c.Writer.Write(raw); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to write json file")
	}
}

func truncate(value string, max int) string {
	runes := []rune(value)
	if len(runes) <= max {
		return value
	}
	return string(runes[:max]) + "..."
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
