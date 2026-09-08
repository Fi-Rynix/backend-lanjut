CREATE TABLE IF NOT EXISTS students (
    id         SERIAL       PRIMARY KEY,
    nim        INTEGER      NOT NULL UNIQUE,
    name       VARCHAR(100) NOT NULL,
    grade      DECIMAL(4,2) NOT NULL,
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_students_is_active
    ON students(is_active);
