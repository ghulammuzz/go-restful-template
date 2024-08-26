# Go Restful Template

Go Restful Template adalah proyek template untuk membangun API RESTful menggunakan bahasa pemrograman Go.

## Struktur Proyek

on-going


## Fitur

- **Middleware**: Termasuk middleware untuk autentikasi JWT, kompresi gzip, CORS, dan rate limiting.
- **Logging**: Terintegrasi dengan `log/slog` untuk logging yang fleksibel dengan level log yang dapat dikonfigurasi.
- **Error Handling**: Pengelolaan error yang terstruktur dengan kode dan pesan error yang dapat dikustomisasi.
- **Service Layer**: Memisahkan logika bisnis dari logika aplikasi untuk pemeliharaan yang lebih baik.
- **Repository Layer**: Mengelola interaksi dengan database menggunakan pola repository.
- **Model Layer**: Menyediakan struktur data untuk aplikasi.
- **Configuration**: Mendukung konfigurasi database dan lingkungan melalui file env dan skrip SQL.
- **Deployment**: Mendukung build dan deploy menggunakan Docker dan Terraform.

## Instalasi

### Persyaratan

- Go 1.18 atau versi lebih baru
- Docker (untuk build dan deploy)
- Terraform (untuk manajemen infrastruktur)

### Langkah-langkah

1. **Clone Repositori**: 
   ```bash
   git clone https://github.com/username/go-restful-template.git
2. Instal Dependensi:

    ```bash
    Copy code
    go mod 
    ```
3. Konfigurasi Lingkungan: Salin file `.env.example` ke `.env` dan sesuaikan dengan konfigurasi Anda:

    ```bash
    Copy code
    cp example.env .env
    ```
4. Build dan Jalankan Aplikasi: Gunakan Docker untuk membangun dan menjalankan aplikasi di lingkungan pengembangan atau produksi:

    ```bash
    Copy code
    docker build -f deploy/docker/dev.Dockerfile -t go-restful-template:dev .
    docker run -p 8080:8080 --env-file .env go-restful-template:dev
    ```
5. Deploy dengan Terraform:

    ```bash
    Copy code
    cd deploy/terraform
    terraform init
    terraform apply
    ```

### Penggunaan
- Endpoints: Dokumentasi endpoint API dapat ditemukan di folder docs.
- Logging: Pengaturan logging dapat dikonfigurasi dengan mengubah variabel `APP_PROFILE` di file `.env`.
- Database: Skrip SQL untuk setup database ada di `config/user.sql`.