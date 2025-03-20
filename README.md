# Basic Golang Framework Mux

Aplikasi sederhana CRUD menggunakan framework Mux dan struktur clean architecture, dikemas didalam Docker.

## Fitur  

- **Enkripsi Password dengan bcrypt**: Semua password pengguna di-enkripsi menggunakan metode bcrypt untuk keamanan.  
- **Error Handler**: Penanganan kesalahan yang efektif untuk setiap permintaan.  
- **Level Log**: Logging untuk informasi, peringatan, dan kesalahan dalam aplikasi yang tersimpan di file server.log.

## Endpoint API 


## Prasyarat  

Pastikan Anda memiliki:  

- [Docker](https://www.docker.com/get-started) terpasang di komputer Anda.  
- [Docker Compose](https://docs.docker.com/compose/install/) terinstal.  

## Cara Menjalankan Aplikasi  

1. **Clone repositori**:  

   ```bash  
   git clone https://github.com/Dstar18/GoMux
   cd GoMux  

2. **Jalankan docker compose**:  

   ```bash  
   docker compose -f 'docker-compose.yml' up -d --build 'postgres'
 
   docker compose -f 'docker-compose.yml' up -d --build 'adminer'

   docker compose -f 'docker-compose.yml' up -d --build 'app'

## Access Aplikasi dan Port  

- Apps: localhost:3000
- PostgreSQL: localhost:5432
- Adminer: localhost:8081