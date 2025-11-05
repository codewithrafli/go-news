package app

import (
	"bwanews/config"

	"github.com/rs/zerolog/log"
)

// Fungsi untuk menjalankan server
func RunServer() {
	// Membuat instance config
	cfg := config.NewConfig()

	// Membuat koneksi ke database
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return
	}

}
