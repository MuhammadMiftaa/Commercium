# Commercium Backend Documentation

## Overview

Project ini adalah aplikasi backend yang dibangun dengan Go, menggunakan PostgreSQL sebagai database, dan framework seperti Gorm dan Gin untuk ORM dan routing. Aplikasi ini memiliki sistem otentikasi berbasis cookie dan mendukung modularisasi berdasarkan entitas dengan menerapkan dependency injection.

## Project Structure

Berikut adalah struktur direktori dan deskripsi fungsi masing-masing komponen utama:

### 1. **Main Program**

- **Path**: `/cmd/main.go`
- **Deskripsi**:
  File ini merupakan entry point untuk aplikasi. Di sini, fungsi utama untuk membuat koneksi ke database PostgreSQL diinisialisasi menggunakan Gorm. Fungsi ini mengembalikan koneksi dengan tipe `*gorm.DB`, yang kemudian diinject ke dalam fungsi router.

### 2. **Router**

- **Path**: `/interface/http/router`
- **Deskripsi**:
  File ini mengelompokkan endpoint aplikasi dan menginisialisasi routing masing-masing entitas. Hasil konfigurasi ini mengembalikan nilai bertipe `*gin.Engine` dengan semua grouping dan routing endpoint.

### 3. **Routes per Entity**

- **Path**: `/interface/http/routes`
- **Deskripsi**:
  Folder ini mengelola routing untuk setiap entitas di aplikasi. Di sini dilakukan dependency injection dan penambahan middleware otentikasi. Setiap endpoint akan memanggil handler yang diambil dari struct handler yang sudah diinject.

### 4. **Authentication Middleware**

- **Path**: `/interface/http/middlewares`
- **Deskripsi**:
  Middleware otentikasi yang memvalidasi cookie yang dihasilkan dari proses login. Pengguna harus memiliki cookie ini untuk mengakses handler terkait.

### 5. **Handlers per Entity**

- **Path**: `/interface/http/handler`
- **Deskripsi**:
  Setiap handler di sini memanggil method di layer service yang telah diinject. Handler ini menangani request dari client, memprosesnya, dan memanggil service untuk mengelola logika bisnis.

### 6. **Service per Entity**

- **Path**: `/internal/service`
- **Deskripsi**:
  Di layer ini, service berinteraksi dengan layer repository untuk menjalankan logika bisnis. Service layer di sini memanggil method dari repository yang sudah diinject.

### 7. **Repository per Entity**

- **Path**: `/internal/repository`
- **Deskripsi**:
  Di sini, repository berinteraksi langsung dengan `*gorm.DB` yang diinject dari layer utama untuk melakukan operasi CRUD pada database PostgreSQL.

---

## How to Run the Project

Untuk menjalankan aplikasi, ikuti langkah-langkah berikut:

1. **Install Dependencies**: Pastikan semua package dan dependencies telah terinstall.
   ```bash
   go mod tidy
   ```
2. **Create Database**: Buatlah database dengan nama `commercium`

3. **Set Environment Variables**: Pastikan file `.env` terkonfigurasi dengan benar dengan membuat variable `DB_USER` dan `DB_PASSWORD` sesuai konfigurasi PostgreSQL.

4. **Run the Application**: Eksekusi main program.
   ```bash
   go run cmd/main.go
   ```

## Contributing

Jika ingin berkontribusi, pastikan untuk membuat branch baru dan lakukan pull request setelah perubahan selesai. Pastikan kode mengikuti pedoman penulisan yang ada.

---
