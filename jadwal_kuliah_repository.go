package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"api-students/app/model"
)

type JadwalKuliahRepository interface {
	FindJadwalByNIM(ctx context.Context, nim string) (model.JadwalKuliah, error)
}

type jadwalKuliahPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewJadwalKuliahRepository(pool *pgxpool.Pool) JadwalKuliahRepository {
	return &jadwalKuliahPostgresRepository{pool: pool}
}

func (r *jadwalKuliahPostgresRepository) FindJadwalByNIM(
	ctx context.Context, nim string,
) (model.JadwalKuliah, error) {
	var j model.JadwalKuliah

		err := r.pool.QueryRow(ctx,
		`SELECT id_kuliah, id_student, mata_kuliah, hari
		FROM jadwal_kuliah left join students s on j.id_student = s.id
		WHERE s.NIM = $1`, nim,
	).Scan(&j.IDKuliah, &j.IDStudent, &j.MataKuliah, &j.Hari)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.JadwalKuliah{}, ErrNotFound
		}
		return model.JadwalKuliah{}, fmt.Errorf("mengambil jadwal kuliah: %w", err)
	}

	return j, nil
}