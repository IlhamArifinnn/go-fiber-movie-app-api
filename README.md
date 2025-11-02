# 🎬 Go Fiber Movie App API — Starter Pack

> StudyJam Web Development 2025 - GDGoC STT-NF

Selamat datang di **Starter Pack Movie App API**!  
Project ini dibuat khusus untuk peserta **StudyJam Web Development 2025** sebagai bahan latihan membuat **RESTful API dengan Go Fiber dan PostgreSQL**.

---

## 🚀 Overview

Aplikasi ini merupakan **API sederhana untuk manajemen data film (CRUD Movies)**.  
Kamu bisa melakukan:

- Menampilkan semua data film 🎥
- Melihat detail film berdasarkan ID 🔍
- Menambahkan data film baru ➕
- Mengedit data film ✏️
- Menghapus data film ❌

Project ini sudah dilengkapi dengan:

- **Fiber Framework** — HTTP web framework cepat untuk Go
- **PostgreSQL** — database relasional modern dan powerful
- **GORM** — ORM (Object Relational Mapper) untuk Go
- **Zerolog** — structured logging cepat dan efisien
- **Validator** — validasi input otomatis
- **Swagger** _(optional)_ — dokumentasi API otomatis
- **Air** — hot reload untuk pengembangan

---

<br />

## 🧠 Prerequisites

Sebelum mulai, pastikan kamu sudah menginstal:

| Tool                                               | Versi Minimum | Keterangan               |
| -------------------------------------------------- | ------------- | ------------------------ |
| [Go](https://go.dev/dl/)                           | 1.25          | Bahasa pemrograman utama |
| [PostgreSQL](https://www.postgresql.org/download/) | 17            | Database yang digunakan  |
| [Air](https://github.com/cosmtrek/air)             | latest        | Hot reload (dev only)    |
| [Git](https://git-scm.com/)                        | -             | Untuk clone repo         |
| [Swagger](https://swagger.io/tools/swagger-ui/)    | (opsional)    | Dokumentasi API otomatis |

---

<br />

## 🧩 Tech Stack

| Kategori      | Teknologi                                                             |
| ------------- | --------------------------------------------------------------------- |
| Framework     | [Go Fiber v2](https://github.com/gofiber/fiber)                       |
| Database      | [PostgreSQL](https://www.postgresql.org/)                             |
| ORM           | [GORM](https://gorm.io/)                                              |
| Logging       | [Zerolog](https://github.com/rs/zerolog)                              |
| Validation    | [Go Playground Validator](https://github.com/go-playground/validator) |
| Config Loader | [Godotenv](https://github.com/joho/godotenv)                          |
| JSON Encoder  | [Sonic](https://github.com/bytedance/sonic)                           |
| Dev Tools     | [Air](https://github.com/cosmtrek/air), Swagger                       |

---

<br />

## 🗂️ Struktur Database

### Tabel: `movies`

| Kolom              | Tipe Data    | Deskripsi               | Validasi                       |
| ------------------ | ------------ | ----------------------- | ------------------------------ |
| `title`            | VARCHAR(255) | Judul film              | required                       |
| `description`      | TEXT         | Deskripsi film          | required                       |
| `poster_url`       | VARCHAR(255) | URL poster film         | required, url                  |
| `release_date`     | DATE         | Tanggal rilis           | required, format: `YYYY-MM-DD` |
| `rating`           | DECIMAL(3,1) | Rating film             | required, numeric              |
| `duration_minutes` | INT          | Durasi film dalam menit | required, numeric              |
| `director`         | VARCHAR(255) | Nama sutradara          | required                       |
| `genre`            | JSON         | List genre film         | required, minimal 1 item       |

🧮 **Contoh Data JSON**

```json
{
  "title": "Inception",
  "description": "A mind-bending thriller about dreams within dreams.",
  "poster_url": "https://example.com/inception.jpg",
  "release_date": "2010-07-16",
  "rating": 8.8,
  "duration_minutes": 148,
  "director": "Christopher Nolan",
  "genre": ["Sci-Fi", "Thriller"]
}
```

---

<br />

## 📦 Instalasi & Setup

> Ikuti langkah-langkah berikut untuk menjalankan aplikasi secara lokal:
> <br />

1. **Clone repository ini**

```bash
  git clone https://github.com/zdacoder go-fiber-movie-app-api.git
```

  <br />

2. **Masuk ke direktori project**

```bash
   cd go-fiber-movie-app-api
```

  <br />

3. **Install dependencies**

```bash
   go mod download
```

  <br />

4. **Buat database PostgreSQL**
   Buat database baru di PostgreSQL, misalnya dengan nama `movie_app_db`.

  <br />

5. **Salin file `.env.example` menjadi `.env`**

```bash
   cp .env.example .env
```

  <br />

6. **Edit file `.env` sesuai konfigurasi database kamu**
   Sesuaikan variabel `DB_NAME`, `DB_USER`, `DB_PASSWORD`, dll. di file `.env`.

  <br />

7. **Jalankan aplikasi**

```bash
   air
```

  <br />

8. **Akses API**
   Buka browser atau tools seperti Postman, dan akses API di:

```
   http://localhost:3000/api/movies
```

---

<br />

## 📄 Dokumentasi API dengan Swagger (Opsional)

Jika kamu ingin menggunakan Swagger untuk dokumentasi API otomatis, ikuti langkah berikut:

1. **Inisialisasi Swagger**
   Pastikan kamu sudah mengimpor package Swagger di `main.go` dan menambahkan route untuk Swagger UI.

2. **Jalankan perintah swag untuk generate dokumentasi**

```bash
   swag init
```

3. **Akses Swagger UI**
   Buka browser dan akses:

```
   http://localhost:3000/swagger
```

---
