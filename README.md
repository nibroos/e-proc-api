<h1 align="center">
    <br>
  Backend e-Procurement
  <br>
</h1>

Asumsi saya, aplikasi ini dibuat untuk vendor yang ingin membuat manajemen barang yang akan dijual, dan juga untuk pembeli yang ingin membeli barang yang dijual oleh vendor, kita sebagai admin mengelola konfirmasi/persetujuan vendor. Waktu terbatas, sampai saat ini beberapa yang sudah saya implementasikan: 
- Dokumentasi postman sudah ada di root folder, file `e-proc.postman_collection.json` dengan response yang tersimpan sebagai bukti.
- Registrasi Vendor & Pembeli menggunakan endpoint `/auth/register`
- CRUD catalog hanya bagian parent, belum implementasi child(item) catalog, namun sudah ada seedernya. Cek endpoint `/catalogs`
- Docker compose untuk development environment bisa digunakan, production belum dicoba. Docker image sudah ada di docker hub.
- CRUD user dengan role permission. cek `/users`
- Semua endpoint menggunakan JWT token, kecuali endpoint register dan login.
- Hanya menggunakan method POST dengan pertimbangan cepat implementasi & debug yang lebih mudah.
- CRUD seeder & migration, gunakan makefile dari luar service untuk menjalankan seeder & migration. cd ke `/service` lalu jalankan `make migrate-up`. untuk seeder, jalankan endpoint `seeders/run`, ini bukan best practice, hanya shortcut saja.
- Semua select menggunakan raw query dengan pertimbangan performa, ORM untuk create/update.
- Untuk deployment di production, sebenarnya sudah ada script Jenkinsfile yang saya buat dan biasa saya gunakan, namun belum saya coba di project ini. silahkan cek log deployment aplikasi saya (e-learning) yang lain ke link ini: [Jenkins](https://deploy.nibros.tech)
- Semua isu yang saya temui sudah saya catat di bawah, jika ada yang ingin ditanyakan, silahkan. 
- Technical question document bisa diakses di [Technical Question](https://docs.google.com/document/d/1cHUQHMkoz9jnRxcDHQaTMZ9YNCsM2vVwbxgJZwwM4es/edit?usp=sharing)

Mohon maaf banyak bagian yang belum lengkap karena keperluan kantor saat ini yang padat, hanya efektif mengerjakan 1 hari. Saya harap reviewer mengerti garis besarnya dan bisa punya gambaran bagaimana saya bekerja dengan potensi yang saya miliki. 

Terima kasih. 

## 🚀 Quick Start
### Developement Environment
On `/` dir, Run `make copy-env`, Modify to suit your environment, focus on these key, you can leave others as it is. The key name is explanatory itself.
```bash
# MAIN APP PORT
GATEWAY_PORT=4050

# DATABASE
POSTGRES_PORT=5432
```

If the make command is not available, copy `.env.example` to `.env` on `/docker` & `/service` dir, then modify the key, then run `docker-compose -f docker/docker-compose-dev.yml up --build -d` on `/` dir.

> Without docker, you need to install [air-verse](https://github.com/air-verse/air) to activate the hot reloading.

### 🐳 Docker :: Container Platform

[Docker](https://docs.docker.com/get-docker/) Install.

- On the root folder, Starts the containers in the background and leaves them running : `docker-compose -f docker/docker-compose-dev.yml up --build -d`
- Stops containers and removes containers, networks, volumes, and images : `docker-compose down`

## 🛎 Available Commands each Service

Change bash directory to each service.
> ${arg} means replace all of it match your args without space
- Run export path : `export PATH="$PATH:$(go env GOPATH)/bin"`
- Create mirgration : `make migrate-create name=${your_migration_name}`
- Run migration : `make migrate-up`
- Stepback migraiton: `make migrate-down`
- Generate proto file, leave the proto args blank if you want to generate all proto file: `make proto ${your-proto.proto}`. If its fail, run this command on specific service. for example, in /service/ run bash `export PATH="$PATH:$(go env GOPATH)/bin"`
- Create seeder : `make seed-create name=${your_seeder_name}`
- Run seeder : `make seed-run file=${your_seeder_name}.sql`

## 💎 The Package Features

<p>
  <img src="https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge&logo=Docker&logoColor=fff" />&nbsp;&nbsp;
  <img src="https://img.shields.io/badge/-NGINX-269539?style=for-the-badge&logo=NGINX&logoColor=fff" />
  <img src="https://img.shields.io/badge/-Go-1185F4?style=for-the-badge&logo=Go&logoColor=fff" />
</p>
<p>
<img src="https://img.shields.io/badge/-PostgreSQL-336791?style=for-the-badge&logo=PostgreSQL&logoColor=fff" />&nbsp;&nbsp;
</p>

## 📔 Notes & Issues

#### dial tcp: lookup postgres: no such host
Change the makefile DB_HOST to `localhost` if run in local env, when running on docker, change it to `postgres`, make sure no space in the value.

#### run multiple seeder in one execution
You can run multiple seeder references in the seeder_controller.go file.

#### error running migration fix migration
Change the 'version' column name on schema_migrations to latest succeed migration, change the 'dirty' column to false, then run the migration again

#### error function gen_salt(unknown) does not exist, postgre extensions
`CREATE EXTENSION IF NOT EXISTS pgcrypto;`

### 📗 API Document
All endpoints stored in  `-.json`

<h1 align="center">
    <br>
  Features
  <br>
</h1>

Feel free to ask if you have any questions or need more details!

