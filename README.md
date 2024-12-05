# go_restapi_sddp

Proyek ini adalah REST API sederhana yang dikembangkan dengan Golang.

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![Go Reference](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)
Gin is a web framework written in Go. It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

Gin's key features are:

Zero allocation router
Speed
Middleware support
Crash-free
JSON validation
Route grouping
Error management
Built-in rendering
Extensible
Proyek ini adalah REST API sederhana yang dikembangkan dengan Golang.

Getting started

## Daftar Isi
1. [Pendahuluan](#pendahuluan)
2. [Fitur](#fitur)
3. [Persyaratan](#persyaratan)
4. [Instalasi](#instalasi)

## Pendahuluan

Proyek ini adalah REST API sederhana yang dikembangkan dengan Golang. Proyek ini bertujuan untuk memberikan panduan praktis dalam membangun REST API dengan struktur kode yang terorganisir dan menggunakan teknologi modern seperti JWT untuk otentikasi.


Proyek ini bertujuan untuk memberikan contoh implementasi REST API menggunakan Golang, lengkap dengan dokumentasi dan konfigurasi.


## Fitur
- **CRUD Operations**: Perform Create, Read, Update, and Delete operations on perfumes.
- **Authentication with JWT**: Secure the API endpoints using JSON Web Tokens (JWT) for authentication.
- **Input Validation**: Ensure all input data is validated for security and correctness.
- **Logging**: Monitor activity and log important events for debugging and auditing purposes.
- **Database Support**: Connect to MySQL for storing data. You can also adapt the system to support PostgreSQL.
- **Folder Structure**: The project follows standard industry practices for folder and file organization.


## Persyaratan
Sebelum memulai, pastikan Anda sudah memiliki:
- Golang versi 1.20 atau lebih baru
<<<<<<< HEAD
- Xampp 
- Git
- mysql server (download di https://dev.mysql.com/downloads/mysql/)
- xampp
- MySQL kl blm punya download di ([sini](https://dev.mysql.com/downloads/installer/))
- Git
- postman kl blm punya download di ([sini](https://www.postman.com/downloads/))

## Instalasi
Ikuti langkah-langkah berikut untuk menginstal proyek:

1. Clone repository:
   ```bash
   git clone https://github.com/fadlyfebros/go_restapi_sddp.git
   
Pastikan Anda telah menginstal beberapa alat berikut:

XAMPP (termasuk MySQL). Jika belum, Anda dapat mengunduhnya di ([sini](https://dev.mysql.com/downloads/installer/)).
MySQL (termasuk service MySQL yang dijalankan). Jika Anda menggunakan XAMPP, MySQL sudah termasuk di dalamnya.
Git untuk meng-clone repository. Jika belum terinstal, Anda dapat mengunduhnya di ([sini](https://git-scm.com/downloads)).
Go (Golang). Pastikan Anda sudah menginstal Golang versi 1.20 atau lebih baru, dapat diunduh di ([https://go.dev/doc/install](sini)).

2. Clone repository
Clone repository proyek ini ke direktori lokal Anda:
git clone https://github.com/fadlyfebros/go_restapi_sddp.git

3. Konfigurasi MySQL
Jalankan MySQL dengan XAMPP atau terminal jika Anda menggunakan MySQL secara terpisah.
Masuk ke MySQL dengan perintah:
mysql -u root
Jika berhasil, Anda akan melihat pesan Query OK, 1 row affected (0.001 sec). Setelah itu, ketik exit untuk keluar dari MySQL.

4. Menjalankan aplikasi
Setelah repository ter-clone, buka folder proyek dan jalankan perintah berikut untuk menginstal dependensi:
go mod tidy
Jalankan aplikasi dengan perintah:
go run main.go

5. Masalah umum
Jika saat menjalankan aplikasi Anda mendapatkan pesan kesalahan seperti berikut:
2024/11/10 11:06:29 D:/Febro/Website/website me/go_restapi_sddp/moduls/setup.go:16 [error] failed to initialize database, got error dial tcp: lookup http://localhost:8080: no such host
2024/11/10 11:06:29 Failed to connect to the database: dial tcp: lookup http://localhost:8080: no such host
exit status 1
Ini kemungkinan disebabkan oleh konfigurasi yang salah pada localhost:8080. Pastikan Anda telah mengonfigurasi koneksi database MySQL di file setup.go dengan alamat yang benar:
Gantilah:

dsn := "root:@tcp(http://localhost:8080)/go_restapi_sddp"

Menjadi:

dsn := "root:@tcp(127.0.0.1:3306)/go_restapi_sddp"

Setelah melakukan perubahan tersebut, coba jalankan kembali aplikasi menggunakan:

go run main.go
