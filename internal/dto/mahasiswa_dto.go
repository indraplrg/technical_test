package dto

// MahasiswaRequest is the payload for creating/updating a Mahasiswa.
type MahasiswaRequest struct {
	Nama         string `json:"nama" binding:"required"`
	Umur         int    `json:"umur" binding:"required,gt=0"`
	NIM          string `json:"nim" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required,date"`
	Alamat       string `json:"alamat" binding:"required"`
	IDJurusan    uint   `json:"id_jurusan" binding:"required"`
}
