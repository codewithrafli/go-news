package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Struct untuk koneksi ke database PostgreSQL
type Postgres struct {
	DB *gorm.DB
}

// Fungsi untuk mengkoneksikan ke database PostgreSQL
func (cfg Config) ConnectionPostgres() (*Postgres, error) {

	// Membuat string koneksi ke database PostgreSQL
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName,
	)

	log.Info().Str("host", cfg.PsqlDB.Host).Int("port", cfg.PsqlDB.Port).Msg("Connecting to database")

	// Membuat koneksi ke database PostgreSQL
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})

	// Jika gagal mengkoneksikan ke database PostgreSQL
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	// Mendapatkan koneksi ke database PostgreSQL
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database connection")
		return nil, err
	}

	// Mengatur jumlah maksimal koneksi terbuka dan idle
	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	log.Info().Msg("Successfully connected to database")

	// Mengembalikan koneksi ke database PostgreSQL
	return &Postgres{DB: db}, nil
}
