# Simple Greeting API with Gin

API sederhana untuk mengelola pesan salam menggunakan framework Gin di Go.

## Deskripsi

API ini menyediakan endpoint dasar untuk:
- Mendapatkan daftar salam
- Mendapatkan salam berdasarkan ID
- Membuat salam baru

Data disimpan dalam memori (tidak persisten) dan akan direset setiap kali server dimulai ulang.

## Fitur

- GET semua salam
- GET salam by ID
- POST membuat salam baru
- Validasi input
- Penanganan error dasar

## API Endpoints

### GET /api/Salam/greetings
Mendapatkan semua salam
### GET /api/Salam/greetings/id
Mendapatkan salam sesuai id 
### POST /api/Salam/greetings
Membuat salam baru
