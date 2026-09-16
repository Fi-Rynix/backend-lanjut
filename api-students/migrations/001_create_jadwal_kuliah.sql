CREATE TABLE IF NOT EXISTS jadwal_kuliah (
    id_kuliah         SERIAL       PRIMARY KEY,
    id_student SERIAL NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    mata_kuliah VARCHAR(100) NOT NULL,
    hari        VARCHAR(10) NOT NULL,
);