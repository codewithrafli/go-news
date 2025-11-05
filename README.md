# BWA News - Portal Berita dengan Golang

Portal berita modern yang dibangun menggunakan Go dengan arsitektur Clean Architecture.

## Deskripsi Project

BWA News adalah aplikasi web portal berita yang dikembangkan dengan bahasa pemrograman Go. Project ini dibuat untuk tujuan pembelajaran dalam membangun aplikasi web yang scalable dan maintainable menggunakan prinsip Clean Architecture.

## Tech Stack

- **Backend Framework:** Go 1.24.4
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Token)
- **Storage:** Cloudflare (R2/Images)
- **Architecture Pattern:** Clean Architecture

## Struktur Project

```
bwanewsgo/
├── cmd/                    # Entry points untuk berbagai command
├── config/                 # Konfigurasi aplikasi
│   └── config.go
├── database/              # Database migrations dan seeds
├── internal/              # Kode internal aplikasi
│   ├── adapter/           # Adapter layer (Framework & Drivers)
│   │   ├── handler/       # HTTP handlers
│   │   ├── repository/    # Database repositories
│   │   └── cloudfare/     # Cloudflare integration
│   ├── core/              # Business logic layer
│   │   ├── domain/        # Domain models/entities
│   │   └── service/       # Business services
│   └── app/               # Application layer
├── lib/                   # Shared libraries
│   ├── conv/              # Conversion utilities
│   └── jwt/               # JWT utilities
├── .env                   # Environment variables
├── go.mod                 # Go modules
└── main.go               # Main application entry point
```

## Prinsip Clean Architecture

Project ini menggunakan Clean Architecture dengan lapisan-lapisan berikut:

1. **Domain Layer** (`internal/core/domain`): Berisi entities dan business rules
2. **Service Layer** (`internal/core/service`): Berisi use cases dan business logic
3. **Repository Layer** (`internal/adapter/repository`): Berisi implementasi database
4. **Handler Layer** (`internal/adapter/handler`): Berisi HTTP handlers dan routing
5. **Application Layer** (`internal/app`): Orchestration dan dependency injection

## Prasyarat

Sebelum menjalankan project ini, pastikan Anda telah menginstall:

- Go 1.24.4 atau lebih baru
- PostgreSQL
- Git

## Instalasi

1. Clone repository ini:
```bash
git clone <repository-url>
cd bwanewsgo
```

2. Install dependencies:
```bash
go mod download
```

3. Setup environment variables:
```bash
cp .env.example .env
```

4. Konfigurasi file `.env` dengan kredensial Anda:
```env
APP_ENV="development"
APP_PORT="8080"

DATABASE_PORT=5432
DATABASE_HOST=localhost
DATABASE_USER=your_db_user
DATABASE_PASSWORD=your_db_password
DATABASE_NAME=bwanews
DATABASE_MAX_OPEN_CONNECTION=10
DATABASE_MAX_IDLE_CONNECTION=10

JWT_SECRET_KEY=your_secret_key_here
JWT_ISSUER=bwanews
```

5. Jalankan database migrations (coming soon)

## Menjalankan Aplikasi

### Development Mode

```bash
go run main.go
```

### Production Build

```bash
go build -o bwanews main.go
./bwanews
```

## Fitur (Planned/In Development)

- [ ] User Authentication (Register, Login, Logout)
- [ ] User Authorization dengan JWT
- [ ] CRUD Artikel/Berita
- [ ] Kategori Berita
- [ ] Upload Gambar ke Cloudflare
- [ ] Pencarian Berita
- [ ] Komentar pada Berita
- [ ] User Profile Management
- [ ] Admin Dashboard

## API Documentation

API documentation akan tersedia setelah implementasi handler selesai.

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Contributing

Project ini dibuat untuk tujuan pembelajaran. Jika Anda ingin berkontribusi:

1. Fork repository ini
2. Buat feature branch (`git checkout -b feature/amazing-feature`)
3. Commit perubahan Anda (`git commit -m 'Add some amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buat Pull Request

## Learning Resources

Jika Anda baru belajar Go dan Clean Architecture, berikut beberapa resource yang bermanfaat:

- [Go Documentation](https://go.dev/doc/)
- [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Clean Architecture Example](https://github.com/bxcodec/go-clean-arch)

## License

Project ini dibuat untuk tujuan pembelajaran.

## Kontak

Jika ada pertanyaan, silakan buat issue di repository ini.

---

**Happy Learning! 🚀**
