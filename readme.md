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
2. Buat src/infra semua keperluan infrastruktur
3. Buat src/infra/config, lalu buat file config.go untuk setting semua configurasinya
4. Buat src/infra/constants dan jangan lupa lakukan go mod tidy untuk mengintegrasikan package yang dibutuhkan
5. Buat src/infra/errors untuk custom error messages
6. Buat src/infra/helper untuk nantinya meletakkan semua helper yang dibutuhkan. untuk case awal ini disitu berisi keperluan untuk JWT
7. Buat src/infra/log untuk custom log
8. Buat src/infra/persistence untuk keperluan persistensi data
9. Buat src/infra/persistence/postgres/postgres.go untuk configurasi postgresql. jangan lupa go mod tidy
10. Buat  src/infra/persistence/redis/redis.go untuk configurasi redis.
11. Buat  ../infra/broker/nats/nats.go, untuk configurasi nats

# STEP 2

## URUTAN TAHAP 2
1. Buat direktori app untuk semua kebutuhan app nantinya, dan buat direktori dto untuk semua kebutuhan data transfer object nantinya
2. pada direktori dto, buat ../user/user.go.
3. pada direktori app, buat ../repositories/user/user.go.
4. buat direktori usecases pada direktori app. lalu buat direktori user pada direktori usecases.
5. pada direktori usecases buat file usecases.go untuk warping semua usecases nantinya.
5. buat file interface/rest/response/response.go
6. buat direktori src/infra/interface/rest, lalu buat direktori ../interface/handler, dan ../interface/route
7. buat ../handler/user/user.go dan ../route/user.go
8. buat ../interface/rest/rest.go untuk method yang mengintegrasikan routing
9. buat main.go untuk configurasi semuanya





