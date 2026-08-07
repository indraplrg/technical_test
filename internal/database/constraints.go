package database

import "gorm.io/gorm"

// foreignKeyConstraints contains idempotent SQL statements that create the
// foreign key constraints between mahasiswa and jurusan. They are applied
// after GORM AutoMigrate, which intentionally does not manage FK constraints.
const foreignKeyConstraints = `
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_mahasiswa_jurusan') THEN
        ALTER TABLE mahasiswa
            ADD CONSTRAINT fk_mahasiswa_jurusan
            FOREIGN KEY (id_jurusan) REFERENCES jurusan (id_jurusan)
            ON UPDATE CASCADE ON DELETE RESTRICT;
    END IF;
END $$;
`

// ApplyConstraints ensures the database foreign keys exist.
func ApplyConstraints(db *gorm.DB) error {
	return db.Exec(foreignKeyConstraints).Error
}
