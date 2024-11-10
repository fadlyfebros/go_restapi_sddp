# go_restapi_sddp

Proyek ini adalah REST API sederhana yang dikembangkan dengan Golang.

## Daftar Isi
1. [Pendahuluan](#pendahuluan)
2. [Fitur](#fitur)
3. [Persyaratan](#persyaratan)
4. [Instalasi](#instalasi)

## Pendahuluan
Proyek ini bertujuan untuk memberikan contoh implementasi REST API menggunakan Golang, lengkap dengan dokumentasi dan konfigurasi.

## Fitur
- CRUD operasi (Create, Read, Update, Delete)
- Otentikasi menggunakan JWT
- Validasi input
- Struktur folder yang terorganisir

## Persyaratan
Sebelum memulai, pastikan Anda sudah memiliki:
- Golang versi 1.20 atau lebih baru
- xampp
- MySQL kl blm punya download di ([https://dev.mysql.com/downloads/installer/](https://dev.mysql.com/downloads/installer/))
- Git

## Instalasi
Ikuti langkah-langkah berikut untuk menginstal proyek:

1. Persiapkan lingkungan pengembangan
Pastikan Anda telah menginstal beberapa alat berikut:

XAMPP (termasuk MySQL). Jika belum, Anda dapat mengunduhnya di ([https://dev.mysql.com/downloads/installer/](sini)).
MySQL (termasuk service MySQL yang dijalankan). Jika Anda menggunakan XAMPP, MySQL sudah termasuk di dalamnya.
Git untuk meng-clone repository. Jika belum terinstal, Anda dapat mengunduhnya di ([https://git-scm.com/downloads](sini)).
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