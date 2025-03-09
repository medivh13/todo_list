# Project Ramadhan TODO LIST 

Project ini dibuat sebagai bagian sharing saya dengan tajuk tantangan belajar coding selama Ramadhan untuk mempelajari microservices, pemanfaatan redis ttl, dan event carrier state transfer.
Aplikasi ini adalah **To do List** berbasis Golang dengan **PostgreSQL, Redis, dan NATS**, yang memungkinkan pengguna untuk mencatat tugas, memantau progress, dan mengatur waktu kadaluarsa tugas secara otomatis.  

## **Tech Stack yang Digunakan**  
- **Golang** → Bahasa pemrograman utama  
- **PostgreSQL** → Database utama untuk menyimpan data tugas  
- **Redis** → Digunakan untuk caching dan TTL tugas  
- **NATS** → Event-driven system untuk komunikasi antar service  
- **JWT** → Digunakan untuk autentikasi pengguna  
- **Docker** → Untuk menjalankan layanan dengan lebih mudah  

# STEP 1

## URUTAN MEMBUAT DARI AWAL
1. Buat database dan create table sesuai pada direktori db
2. Buat direktori src, lalu buat direktori infra pada src untuk semua keperluan infrastruktur
3. Buat direktori config pada direktori infra. lalu buat file config.go untuk setting semua configurasinya
4. Buat direktori constants dan jangan lupa lakukan go mod tidy untuk mengintegrasikan package yang dibutuhkan
5. Buat direktori errors untuk custom error messages
6. Buat direktori helper untuk nantinya meletakkan semua helper yang dibutuhkan. untuk case awal ini disitu berisi keperluan untuk JWT
7. Buat direktori log untuk custom log
8. Buat direktori persistence untuk keperluan persistensi data
9. Buat direktori postgres pada direktori persistence dan buat file postgres.go untuk configurasi postgresql. jangan lupa go mod tidy
10. Buat direktori redis pada dorektori persistence dan buat file redis.go untuk configurasi redis.
11. Buat direktori broker pada infra, lalu buat direktori nats pada direktori broker. lalu buat file nats.go untuk configurasi nats


