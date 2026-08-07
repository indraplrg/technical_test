# API Documentation — Student Management System

Dokumentasi lengkap REST API **Student Management System** (Student Management API v1).

- **Base URL (lokal):** `http://localhost:8080`
- **API Versioning:** semua endpoint berada di bawah prefix `/api/v1`
- **Swagger UI:** `http://localhost:8080/swagger/index.html`
- **Format:** JSON (`application/json`)
- **Timezone:** seluruh timestamp menggunakan UTC (`RFC3339`)
- **Autentikasi:** belum diterapkan (public API)

---

## 1. Konvensi Umum

### 1.1 Envelope Respons Sukses

```json
{
  "success": true,
  "message": "string",
  "data": { }
}
```

- `data` bersifat opsional (tidak muncul jika kosong).
- `data` dapat berupa objek, array, atau struktur `items` + `pagination` (endpoint list).

### 1.2 Envelope Respons Error

```json
{
  "success": false,
  "message": "string",
  "errors": [ ]
}
```

- `errors` bersifat opsional, biasanya berisi daftar pesan validasi per field.

### 1.3 Format Data Terpaginasi (List)

```json
{
  "success": true,
  "message": "mahasiswa fetched successfully",
  "data": {
    "items": [ ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 42,
      "total_pages": 5
    }
  }
}
```

Keterangan:

| Field        | Tipe  | Deskripsi                           |
| ------------ | ----- | ----------------------------------- |
| `items`      | array | data pada halaman aktif             |
| `pagination.page` | int | nomor halaman saat ini          |
| `pagination.limit` | int | jumlah data per halaman         |
| `pagination.total` | int64 | total seluruh data           |
| `pagination.total_pages` | int | total halaman            |

### 1.4 HTTP Status Codes

| Code | Makna                                             | Contoh situasi                              |
| ---- | ------------------------------------------------- | ------------------------------------------- |
| 200  | OK                                                | Get, Update, Delete sukses                  |
| 201  | Created                                           | Create sukses                               |
| 204  | No Content                                        | Respons preflight CORS (OPTIONS)            |
| 400  | Bad Request                                       | ID path invalid, validasi binding gagal     |
| 404  | Not Found                                         | Data tidak ditemukan                        |
| 409  | Conflict                                          | NIM / nama jurusan duplikat, data masih dipakai |
| 500  | Internal Server Error                             | Error tak terduga di server                 |

### 1.5 Header Khusus

- `X-Request-ID` — dihasilkan otomatis (UUID) bila tidak dikirim klien, dipakai untuk tracing log.
- `Access-Control-Allow-Origin` — diset dari `CORS_ALLOWED_ORIGIN` (default `*`).

---

## 2. Health Check

### `GET /health`

Mengetahui status layanan.

**Respons 200:**

```json
{
  "success": true,
  "message": "service healthy",
  "data": { "status": "ok" }
}
```

---

## 3. Jurusan (Departemen)

Base path: `/api/v1/jurusan`

### 3.1 Create Jurusan

`POST /api/v1/jurusan`

**Request Body:**

```json
{
  "nama_jurusan": "Teknik Informatika",
  "fakultas": "Fakultas Ilmu Komputer",
  "jenjang": "S1"
}
```

| Field          | Tipe   | Wajib | Keterangan            |
| -------------- | ------ | ----- | --------------------- |
| `nama_jurusan` | string | ya    | nama jurusan (unik)   |
| `fakultas`     | string | ya    | nama fakultas         |
| `jenjang`      | string | ya    | jenjang (S1, S2, D3)  |

**Respons 201 Created:**

```json
{
  "success": true,
  "message": "jurusan created",
  "data": {
    "id_jurusan": 6,
    "nama_jurusan": "Teknik Informatika",
    "fakultas": "Fakultas Ilmu Komputer",
    "jenjang": "S1",
    "created_at": "2026-08-07T09:07:56.134068+08:00",
    "updated_at": "2026-08-07T09:07:56.134068+08:00"
  }
}
```

**Respons error:**

| Code | Kondisi                          | Contoh `message`                    |
| ---- | -------------------------------- | ----------------------------------- |
| 400  | field wajib kosong               | `validation error` + `errors` array |
| 409  | `nama_jurusan` sudah dipakai     | `nama jurusan already exists`       |
| 500  | error internal                   | `internal server error`             |

---

### 3.2 Get All Jurusan

`GET /api/v1/jurusan`

Mengembalikan seluruh jurusan (tanpa pagination) urut `nama_jurusan` ASC.

**Respons 200:**

```json
{
  "success": true,
  "message": "jurusan fetched successfully",
  "data": [
    {
      "id_jurusan": 4,
      "nama_jurusan": "Akuntansi",
      "fakultas": "Fakultas Ekonomi",
      "jenjang": "D3",
      "created_at": "2026-08-07T09:07:56.134068+08:00",
      "updated_at": "2026-08-07T09:07:56.134068+08:00"
    },
    {
      "id_jurusan": 1,
      "nama_jurusan": "Teknik Informatika",
      "fakultas": "Fakultas Ilmu Komputer",
      "jenjang": "S1",
      "created_at": "2026-08-07T09:07:56.134068+08:00",
      "updated_at": "2026-08-07T09:07:56.134068+08:00"
    }
  ]
}
```

---

### 3.3 Get Jurusan by ID

`GET /api/v1/jurusan/{id}`

| Param | Tipe | Wajib | Keterangan  |
| ----- | ---- | ----- | ----------- |
| `id`  | int  | ya    | `id_jurusan` |

**Respons 200:** sama seperti objek jurusan di atas.

**Respons error:**

| Code | Kondisi                     | Contoh `message`         |
| ---- | --------------------------- | ------------------------ |
| 400  | `id` bukan angka            | `invalid id parameter`   |
| 404  | jurusan tidak ditemukan     | `jurusan not found`      |

---

### 3.4 Update Jurusan

`PUT /api/v1/jurusan/{id}`

Body sama dengan Create (semua field wajib).

**Respons 200:**

```json
{
  "success": true,
  "message": "jurusan updated successfully",
  "data": { "id_jurusan": 6, "...": "..." }
}
```

**Respons error:**

| Code | Kondisi                        | Contoh `message`                   |
| ---- | ------------------------------ | ---------------------------------- |
| 400  | field wajib kosong / id invalid | `validation error` / `invalid id parameter` |
| 404  | jurusan tidak ditemukan        | `jurusan not found`                |
| 409  | nama jurusan dipakai jurusan lain | `nama jurusan already exists`   |

---

### 3.5 Delete Jurusan

`DELETE /api/v1/jurusan/{id}`

Melakukan **soft delete** (mengisi `deleted_at`).

**Respons 200:**

```json
{ "success": true, "message": "jurusan deleted successfully" }
```

**Respons error:**

| Code | Kondisi                              | Contoh `message`                             |
| ---- | ------------------------------------ | -------------------------------------------- |
| 400  | `id` bukan angka                     | `invalid id parameter`                       |
| 404  | jurusan tidak ditemukan              | `jurusan not found`                          |
| 409  | masih ada mahasiswa yang menaut      | `jurusan still has associated mahasiswa`     |

---

## 4. Mahasiswa (Mahasiswa/Siswa)

Base path: `/api/v1/mahasiswa`

### 4.1 Create Mahasiswa

`POST /api/v1/mahasiswa`

**Request Body:**

```json
{
  "nama": "Budi Santoso",
  "umur": 21,
  "nim": "TI-2026-001",
  "tanggal_lahir": "2005-03-15",
  "alamat": "Jl. Merdeka 1, Jakarta",
  "id_jurusan": 1
}
```

| Field            | Tipe   | Wajib | Keterangan                                  |
| ---------------- | ------ | ----- | ------------------------------------------- |
| `nama`           | string | ya    | nama mahasiswa                              |
| `umur`           | int    | ya    | harus > 0                                   |
| `nim`            | string | ya    | NIM unik                                    |
| `tanggal_lahir`  | string | ya    | format `YYYY-MM-DD`                         |
| `alamat`         | string | ya    | alamat                                      |
| `id_jurusan`     | int    | ya    | harus mengacu pada jurusan yang ada         |

**Respons 201 Created** (menyertakan objek jurusan terkait):

```json
{
  "success": true,
  "message": "mahasiswa created",
  "data": {
    "id": 1,
    "nama": "Budi Santoso",
    "umur": 21,
    "nim": "TI-2026-001",
    "tanggal_lahir": "2005-03-15",
    "alamat": "Jl. Merdeka 1, Jakarta",
    "id_jurusan": 1,
    "jurusan": {
      "id_jurusan": 1,
      "nama_jurusan": "Teknik Informatika",
      "fakultas": "Fakultas Ilmu Komputer",
      "jenjang": "S1",
      "created_at": "2026-08-07T09:07:56.134068+08:00",
      "updated_at": "2026-08-07T09:07:56.134068+08:00"
    },
    "created_at": "2026-08-07T09:08:11.450517+08:00",
    "updated_at": "2026-08-07T09:08:11.450517+08:00"
  }
}
```

**Respons error:**

| Code | Kondisi                              | Contoh `message`                  |
| ---- | ------------------------------------ | --------------------------------- |
| 400  | validasi gagal / format tanggal salah | `validation error` + `errors` array |
| 404  | `id_jurusan` tidak ada               | `jurusan not found`               |
| 409  | `nim` sudah dipakai                  | `nim already exists`              |

Contoh body validasi gagal:

```json
{
  "success": false,
  "message": "validation error",
  "errors": [
    "nama is required",
    "umur is required",
    "tanggal_lahir must use YYYY-MM-DD format"
  ]
}
```

---

### 4.2 Get All Mahasiswa (List, Search, Filter, Sort, Pagination)

`GET /api/v1/mahasiswa`

**Query Parameters:**

| Param        | Tipe   | Wajib | Default | Keterangan                                        |
| ------------ | ------ | ----- | ------- | ------------------------------------------------- |
| `search`     | string | tidak | -       | pencarian sebagian (`ILIKE`) pada `nama` atau `nim` |
| `nim`        | string | tidak | -       | filter NIM persis (`=`)                            |
| `id_jurusan` | int    | tidak | -       | filter berdasarkan jurusan                        |
| `sort_by`    | string | tidak | `created_at` | kolom sorting: `nama`, `umur`, `nim`, `tanggal_lahir`, `created_at` (whitelist) |
| `sort_order` | string | tidak | `desc`  | `asc` atau `desc`                                 |
| `page`       | int    | tidak | 1       | halaman aktif                                     |
| `limit`      | int    | tidak | 10      | data per halaman, maksimal 100                    |

Contoh request:

```
GET /api/v1/mahasiswa?search=budi&sort_by=created_at&sort_order=desc&page=1&limit=10
GET /api/v1/mahasiswa?id_jurusan=2&limit=5
GET /api/v1/mahasiswa?nim=TI-2026-001
```

**Respons 200:** format terpaginasi (lihat bagian 1.3). Setiap item berbentuk objek mahasiswa lengkap dengan relasi `jurusan`.

---

### 4.3 Get Mahasiswa by ID

`GET /api/v1/mahasiswa/{id}`

| Param | Tipe | Wajib | Keterangan    |
| ----- | ---- | ----- | ------------- |
| `id`  | int  | ya    | `id` mahasiswa |

**Respons 200:** objek mahasiswa lengkap (termasuk `jurusan`).

**Respons error:**

| Code | Kondisi                   | Contoh `message`       |
| ---- | ------------------------- | ---------------------- |
| 400  | `id` bukan angka          | `invalid id parameter` |
| 404  | mahasiswa tidak ditemukan | `mahasiswa not found`  |

---

### 4.4 Update Mahasiswa

`PUT /api/v1/mahasiswa/{id}`

Body sama dengan Create (semua field wajib). NIM tetap boleh memakai nilai yang sama dengan record itu sendiri (bukan duplikat).

**Respons 200:**

```json
{
  "success": true,
  "message": "mahasiswa updated successfully",
  "data": { "id": 1, "...": "..." }
}
```

**Respons error:**

| Code | Kondisi                              | Contoh `message`                  |
| ---- | ------------------------------------ | --------------------------------- |
| 400  | validasi gagal / id invalid          | `validation error` / `invalid id parameter` |
| 404  | mahasiswa atau jurusan tidak ada     | `mahasiswa not found` / `jurusan not found` |
| 409  | `nim` dipakai mahasiswa lain         | `nim already exists`              |

---

### 4.5 Delete Mahasiswa

`DELETE /api/v1/mahasiswa/{id}`

Melakukan **soft delete** (mengisi `deleted_at`). Record yang sudah dihapus tidak muncul di list/get, dan NIM-nya dapat dipakai kembali.

**Respons 200:**

```json
{ "success": true, "message": "mahasiswa deleted successfully" }
```

**Respons error:**

| Code | Kondisi                   | Contoh `message`       |
| ---- | ------------------------- | ---------------------- |
| 400  | `id` bukan angka          | `invalid id parameter` |
| 404  | mahasiswa tidak ditemukan | `mahasiswa not found`  |

---

## 5. Export Mahasiswa

Semua endpoint export menerima query parameter filter yang sama: `search`, `nim`, `id_jurusan`. Data yang diexport mengikuti filter tersebut (tanpa pagination), urut `nama` ASC.

### 5.1 Export CSV

`GET /api/v1/mahasiswa/export/csv`

- **Content-Type:** `text/csv; charset=utf-8`
- **Content-Disposition:** `attachment; filename=mahasiswa_<timestamp>.csv`

Kolom: `ID, Nama, Umur, NIM, TanggalLahir, Alamat, IDJurusan, Jurusan`

```csv
ID,Nama,Umur,NIM,TanggalLahir,Alamat,IDJurusan,Jurusan
1,Budi Santoso,21,TI-2026-001,2005-03-15,Jl. Merdeka 1 Jakarta,1,Teknik Informatika
```

### 5.2 Export Excel

`GET /api/v1/mahasiswa/export/excel`

- **Content-Type:** `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- **Content-Disposition:** `attachment; filename=mahasiswa_<timestamp>.xlsx`

Sheet `Mahasiswa` berisi header: `ID, Nama, Umur, NIM, Tanggal Lahir, Alamat, ID Jurusan, Jurusan`.

### 5.3 Export PDF

`GET /api/v1/mahasiswa/export/pdf`

- **Content-Type:** `application/pdf`
- **Content-Disposition:** `attachment; filename=mahasiswa_<timestamp>.pdf`

Format A4 landscape dengan tabel: `ID, Nama, Umur, NIM, Tgl Lahir, Alamat, Jurusan` (alamat dipotong maksimal 70 karakter).

### 5.4 Export JSON

`GET /api/v1/mahasiswa/export/json`

- **Content-Type:** `application/json; charset=utf-8`
- **Content-Disposition:** `attachment; filename=mahasiswa_<timestamp>.json`

Isi file mengikuti envelope respons sukses:

```json
{
  "success": true,
  "message": "mahasiswa exported successfully",
  "data": [
    {
      "id": 1,
      "nama": "Budi Santoso",
      "umur": 21,
      "nim": "TI-2026-001",
      "tanggal_lahir": "2005-03-15",
      "alamat": "Jl. Merdeka 1, Jakarta",
      "id_jurusan": 1,
      "jurusan": { "id_jurusan": 1, "nama_jurusan": "Teknik Informatika" },
      "created_at": "2026-08-07T09:08:11.450517+08:00",
      "updated_at": "2026-08-07T09:08:11.450517+08:00"
    }
  ]
}
```

Contoh penggunaan `curl`:

```bash
curl -OJ "http://localhost:8080/api/v1/mahasiswa/export/csv?search=budi"
```

---

## 6. Aturan Validasi

### Mahasiswa

| Field           | Aturan                                  |
| --------------- | --------------------------------------- |
| `nama`          | wajib                                   |
| `umur`          | wajib, > 0                              |
| `nim`           | wajib, unik                             |
| `tanggal_lahir` | wajib, format `YYYY-MM-DD`              |
| `alamat`        | wajib                                   |
| `id_jurusan`    | wajib, jurusan harus ada di database    |

### Jurusan

| Field          | Aturan          |
| -------------- | --------------- |
| `nama_jurusan` | wajib, unik     |
| `fakultas`     | wajib           |
| `jenjang`      | wajib           |

---

## 7. Tabel Referensi Cepat

| Method | Path                              | Fungsi                              |
| ------ | --------------------------------- | ----------------------------------- |
| GET    | `/health`                         | Health check                        |
| POST   | `/api/v1/jurusan`                 | Buat jurusan                        |
| GET    | `/api/v1/jurusan`                 | List jurusan                        |
| GET    | `/api/v1/jurusan/{id}`            | Detail jurusan                      |
| PUT    | `/api/v1/jurusan/{id}`            | Update jurusan                      |
| DELETE | `/api/v1/jurusan/{id}`            | Hapus jurusan (soft)                |
| POST   | `/api/v1/mahasiswa`               | Buat mahasiswa                      |
| GET    | `/api/v1/mahasiswa`               | List mahasiswa (search/filter/sort/pagination) |
| GET    | `/api/v1/mahasiswa/{id}`          | Detail mahasiswa                    |
| PUT    | `/api/v1/mahasiswa/{id}`          | Update mahasiswa                    |
| DELETE | `/api/v1/mahasiswa/{id}`          | Hapus mahasiswa (soft)              |
| GET    | `/api/v1/mahasiswa/export/csv`    | Export CSV                          |
| GET    | `/api/v1/mahasiswa/export/excel`  | Export Excel                        |
| GET    | `/api/v1/mahasiswa/export/pdf`    | Export PDF                          |
| GET    | `/api/v1/mahasiswa/export/json`   | Export JSON                         |
| GET    | `/swagger/index.html`             | Swagger UI                          |

---

## 8. Contoh Alur Penggunaan (curl)

```bash
# 1. Cek kesehatan
curl http://localhost:8080/health

# 2. Buat jurusan
curl -X POST http://localhost:8080/api/v1/jurusan \
  -H "Content-Type: application/json" \
  -d '{"nama_jurusan":"Teknik Sipil","fakultas":"Fakultas Teknik","jenjang":"S1"}'

# 3. Buat mahasiswa
curl -X POST http://localhost:8080/api/v1/mahasiswa \
  -H "Content-Type: application/json" \
  -d '{"nama":"Budi Santoso","umur":21,"nim":"TI-2026-001","tanggal_lahir":"2005-03-15","alamat":"Jl. Merdeka 1, Jakarta","id_jurusan":1}'

# 4. List mahasiswa + pencarian + pagination
curl "http://localhost:8080/api/v1/mahasiswa?search=budi&sort_by=created_at&sort_order=desc&page=1&limit=10"

# 5. Export CSV dengan filter
curl -OJ "http://localhost:8080/api/v1/mahasiswa/export/csv?id_jurusan=1"

# 6. Swagger UI
# Buka http://localhost:8080/swagger/index.html di browser
```
