-- Initial schema for the Student Management System.
-- The application uses GORM AutoMigrate as the source of truth.
-- This file documents the intended relational schema for reference / manual setup.

CREATE TABLE IF NOT EXISTS jurusan (
    id_jurusan   BIGSERIAL PRIMARY KEY,
    nama_jurusan VARCHAR(100) NOT NULL,
    fakultas     VARCHAR(100) NOT NULL,
    jenjang      VARCHAR(50)  NOT NULL,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS mahasiswa (
    id            BIGSERIAL PRIMARY KEY,
    nama          VARCHAR(100) NOT NULL,
    umur          INTEGER      NOT NULL CHECK (umur > 0),
    nim           VARCHAR(20)  NOT NULL UNIQUE,
    tanggal_lahir VARCHAR(20)  NOT NULL,
    alamat        TEXT         NOT NULL,
    id_jurusan    BIGINT       NOT NULL REFERENCES jurusan (id_jurusan) ON UPDATE CASCADE ON DELETE RESTRICT,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_mahasiswa_id_jurusan ON mahasiswa (id_jurusan);
CREATE INDEX IF NOT EXISTS idx_jurusan_deleted_at ON jurusan (deleted_at);
CREATE INDEX IF NOT EXISTS idx_mahasiswa_deleted_at ON mahasiswa (deleted_at);