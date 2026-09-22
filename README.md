# GoAPI Sekolah — Gin + GORM + SQLite

API sederhana untuk data sekolah (Kelas, Siswa, Kartu Pelajar).
Database memakai file `data.sqlite` yang sudah ada di folder project.

## 1. Syarat

- Go 1.27.1+ (project ini diinstall via Homebrew di macOS)
- File `data.sqlite` harus ada di root project

Cek Go:

```bash
go version
```

Kalau di Mac pakai Homebrew dan `go` tidak ketemu:

```bash
export PATH="/opt/homebrew/bin:$PATH"
go version
```

## 2. Install dependency

Cukup sekali (atau setiap ganti dependency):

```bash
go mod tidy
```

Dependency utama:

- `github.com/gin-gonic/gin` — web framework
- `gorm.io/gorm` — ORM
- `gorm.io/driver/sqlite` — driver SQLite

## 3. Menjalankan server

Dari folder project (`goapi/`):

```bash
go run .
```

Kalau berhasil akan muncul:

```
[GIN-debug] Listening and serving HTTP on :8080
```

Buka di browser / curl:

```bash
curl http://localhost:8080/
```

Balasan:

```json
{"message":"API Sekolah jalan! Coba /api/kelas, /api/siswa, /api/kartu"}
```

Server jalan di `http://localhost:8080`.

> Catatan: koneksi database langsung ke `data.sqlite` di folder yang sama.
> Tidak ada AutoMigrate — tabel dianggap sudah ada (`kelas`, `siswas`, `kartu_pelajars`).

## 4. Build binary (opsional)

```bash
go build -o goapi .
./goapi
```

## 5. Daftar endpoint

### Kelas

| Method | Endpoint         | Body                      |
| ------ | ---------------- | ------------------------- |
| GET    | `/api/kelas`     | -                         |
| GET    | `/api/kelas/:id` | -                         |
| POST   | `/api/kelas`     | `{"nama_kelas":"X RPL 2"}` |
| PUT    | `/api/kelas/:id` | `{"nama_kelas":"X RPL 2"}` |
| DELETE | `/api/kelas/:id` | -                         |

### Siswa

| Method | Endpoint         | Body                              |
| ------ | ---------------- | --------------------------------- |
| GET    | `/api/siswa`     | -                                 |
| GET    | `/api/siswa/:id` | -                                 |
| POST   | `/api/siswa`     | `{"nama":"Budi","id_kelas":1}`     |
| PUT    | `/api/siswa/:id` | `{"nama":"Budi","id_kelas":1}`     |
| DELETE | `/api/siswa/:id` | -                                 |

### Kartu Pelajar

| Method | Endpoint         | Body                                            |
| ------ | ---------------- | ----------------------------------------------- |
| GET    | `/api/kartu`     | -                                               |
| GET    | `/api/kartu/:id` | -                                               |
| POST   | `/api/kartu`     | `{"nomor_kartu":"KP-2026-000001","id_siswa":1}` |
| PUT    | `/api/kartu/:id` | `{"nomor_kartu":"KP-2026-000002"}`              |
| DELETE | `/api/kartu/:id` | -                                               |

## 6. Contoh curl

Ambil semua kelas (beserta siswanya):

```bash
curl http://localhost:8080/api/kelas
```

Ambil 1 siswa (beserta kelas + kartunya):

```bash
curl http://localhost:8080/api/siswa/1
```

Tambah kelas:

```bash
curl -X POST http://localhost:8080/api/kelas \
  -H "Content-Type: application/json" \
  -d '{"nama_kelas":"X RPL 2"}'
```

Tambah siswa:

```bash
curl -X POST http://localhost:8080/api/siswa \
  -H "Content-Type: application/json" \
  -d '{"nama":"Budi","id_kelas":1}'
```

Update siswa:

```bash
curl -X PUT http://localhost:8080/api/siswa/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Budi Baru"}'
```

Hapus siswa:

```bash
curl -X DELETE http://localhost:8080/api/siswa/1
```

## 7. Struktur folder

```
goapi/
├── main.go                   → routes + r.Run(":8080")
├── data.sqlite               → database SQLite
├── go.mod / go.sum           → dependency Go
├── config/
│   └── database.go           → konek GORM ke data.sqlite
├── models/
│   └── models.go             → struct Kelas, Siswa, KartuPelajar
└── controllers/
    ├── kelas_controller.go   → CRUD kelas
    ├── siswa_controller.go   → CRUD siswa
    └── kartu_controller.go   → CRUD kartu pelajar
```

## 8. Troubleshooting

- `go: command not found` → Go belum di PATH. Di Mac Homebrew: `export PATH="/opt/homebrew/bin:$PATH"`.
- `Failed to connect to database` → pastikan `data.sqlite` ada di folder yang sama saat run `go run .`.
- Port 8080 dipakai → ganti di `main.go`: `r.Run(":8081")`, lalu akses `http://localhost:8081`.
