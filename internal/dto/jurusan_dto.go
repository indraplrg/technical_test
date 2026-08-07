package dto

// JurusanRequest is the payload for creating/updating a Jurusan.
type JurusanRequest struct {
	NamaJurusan string `json:"nama_jurusan" binding:"required"`
	Fakultas    string `json:"fakultas" binding:"required"`
	Jenjang     string `json:"jenjang" binding:"required"`
}
