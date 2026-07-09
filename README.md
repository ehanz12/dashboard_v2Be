# BE Dashboard

BE Dashboard adalah backend API untuk aplikasi dashboard personal yang digunakan untuk mengelola keuangan, kebiasaan, tugas, jadwal, serta notifikasi reminder. Project ini dibangun menggunakan Go dengan framework Fiber, GORM, MySQL, dan Firebase.

## Fitur Utama

- Autentikasi pengguna
  - Register
  - Login
  - Verifikasi email
  - Lupa/reset password
  - Login dengan Google (ID token)
  - Profil pengguna dan ubah password
- Manajemen transaksi
  - Create, read, update, delete transaksi
- Manajemen kategori
- Manajemen habit / kebiasaan
  - Catatan kebiasaan harian
  - Reminder otomatis
- Manajemen task / tugas
- Manajemen time block / jadwal aktivitas
- Manajemen schedule
- Dashboard & analytics
- Notifikasi user device
- Cron job reminder untuk habit

## Teknologi yang Digunakan

- Bahasa pemrograman: Go
- Web framework: Fiber v2
- ORM: GORM
- Database: MySQL
- Autentikasi: JWT
- Firebase: untuk integrasi layanan Firebase
- Scheduler: gocron
- Environment management: godotenv
- Dokumentasi API: Postman collection

## Struktur Folder

- `config/` - konfigurasi environment dan setting aplikasi
- `cron/` - scheduler reminder habit
- `database/` - koneksi DB dan migrasi SQL
- `dto/` - request/response payload
- `handlers/` - handler HTTP untuk tiap module
- `mappers/` - mapping data antara model dan response
- `middleware/` - middleware autentikasi dan keamanan
- `models/` - model domain aplikasi
- `routers/` - definisi routing API
- `services/` - logika bisnis dan integrasi eksternal
- `utils/` - helper utilitas
- `main.go` - entry point aplikasi

## Prasyarat

Pastikan environment berikut sudah tersedia:

- Go 1.25+
- MySQL
- Firebase service account JSON
- Env file `.env`

## Konfigurasi Environment

Buat file `.env` berdasarkan `.env.example`:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=bedashboard
JWT_SECRET=your_secret
PORT=3040
GOOGLE_CLIENT_ID=your_google_client_id
```

Tambahan konfigurasi email juga dapat diisi jika fitur email aktif.

## Menjalankan Aplikasi

1. Install dependency:

```bash
go mod tidy
```

2. Jalankan migrasi database.

Database migration berada di folder `database/migrations/` dan dapat dipakai sesuai kebutuhan project Anda.

3. Jalankan aplikasi:

```bash
go run .
```

Server akan berjalan di port yang ditentukan pada variabel `PORT`.

## Database

Project ini menggunakan MySQL dan GORM. File SQL awal tersedia di:

- `data.sql`
- `database/migrations/`

## Dokumentasi API

File koleksi Postman tersedia di:

- `postman_collection.json`

Anda dapat mengimpornya ke Postman untuk menguji endpoint API.

## Catatan Tambahan

- Aplikasi ini sudah dilengkapi dengan rate limiting pada API group.
- Terdapat cron job untuk mengirim reminder habit secara berkala.
- CORS sudah dikonfigurasi untuk frontend lokal dan domain tertentu.

## Developer Notes

Struktur project ini mengikuti pola arsitektur modular dengan pemisahan:

- `handlers` untuk HTTP layer
- `services` untuk business logic
- `models` untuk data layer
- `routers` untuk route grouping
- `dto` untuk request/response contract

Hal ini memudahkan pengembangan dan maintenance aplikasi di masa depan.
