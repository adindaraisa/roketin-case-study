# roketin-case-study

Repository ini berisi dua challenge

## Challenge 1: Konversi Waktu Bumi ke Planet Roketin
Program Go sederhana untuk mengonversi waktu dari Bumi ke Planet Roketin.

**Cara menjalankan:**
```bash
go run ./challenge-satu
```

Contoh output:
<p align="center"> <img src="https://github.com/user-attachments/assets/73bb771e-7212-4284-8d5e-1a59ec06318d" alt="Contoh Output Challenge 1" style="max-width: 100%; height: auto;"> </p> ```

## Challenge 2: REST API Movie Festival

Aplikasi REST API untuk mengelola data film festival.

### Fitur:
- Create & update movie
- List semua movie dengan pagination
- Search movie berdasarkan title, description, artists, dan genres
- Versi **tanpa database** (in-memory)
- Versi **dengan database** (PostgreSQL + GORM, tabel ternormalisasi)

---

### Setup Versi Database:
1. Buat database PostgreSQL dan sesuaikan file `.env` dengan contoh di `.env.example`.
2. Import schema dari file `schema.sql` .
3. Jalankan perintah berikut:
```bash
go run ./challenge-dua
go run ./challenge-dua-with-database
```

### Endpoint Utama

| Method | Endpoint         | Deskripsi                                   |
|--------|------------------|---------------------------------------------|
| GET    | `/movies`        | Menampilkan semua film (dengan pagination)  |
| POST   | `/movies`        | Menambahkan film baru                       |
| PUT    | `/movies/{id}`   | Memperbarui data film                       |
| GET    | `/movies/search` | Mencari film berdasarkan query tertentu     |

### Contoh Request (Create Movie)

**POST** `/movies`
```json
{
  "title": "Interstellar",
  "description": "A team of explorers travel through a wormhole in space in an attempt to ensure humanity's survival.",
  "duration": 169,
  "artists": ["Matthew McConaughey", "Anne Hathaway", "Jessica Chastain", "Michael Caine"],
  "genres": ["Adventure", "Drama", "Sci-Fi"]
}
```

### Postman Collection
Postman collection untuk menguji API tersedia di folder postman/Roketin - Case Study.postman_collection.json.
