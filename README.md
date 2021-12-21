# File Service
Aplikasi/service ini digunakan untuk handle upload data ke cloud storage (untuk saat ini firebase storage). Aplikasi/service ini menggunakan protocol HTTP untuk akses datanya.

## Features
- Read Object/file yg terupload di cloud storage
- Upload file khusunya image dan document (pdf,dock) ke cloud storage
- Delete file/object yg sudah di upload di cloud storage

## Technology Used
Service/Aplikasi ini menggunakan bahasa pemrograman Go dan beberap library atau package:
- MySQL
- GoKit 
- Goose 
- Gorilla Mux
- Firebase storage

## How To Run
Jika anda ingin menjalankan aplikasi ini (local development), bisa ikuti langkah-langkah berikut:

- Clone repository ini
 https://github.com/aditya37/file-service.git
- Set beberapa env berikut
    > FIREBASE_BUCKET_NAME = "firebase_bucket_name"
    
    > FIREBASE_CRED_FILE_PATH = "firebase_key.json"
    
    > FIREBASE_PROJECT_ID= "gcp_project_id"
    
    > SERVICE_PORT= 4444
    
    > MYSQL_DBHOST= 127.0.0.1
    
    > MYSQL_DBPORT=3306
    
    > MYSQL_DBUSER= root
    
    > MYSQL_DBPASSWORD= admin
    
    > MYSQL_DBNAME=db_file_service
    
    > MYSQL_MAX_CONNECTION=1
    
    > MYSQL_MAX_IDLE_CONNECTION=1
    
- Kemudian ketika perintah berikut
    > go run main.go


## Endpoint 
Endpoint untuk akses service/aplikasi:
Base Url = http://127.0.0.1:<servie_port>
|  Endpoint | Method  |  Payload | Description|
|-----------|---------|----------|-------------------|
| {Base Url}/file/{object_name}  | GET     | object/file name  |Get detail object by object name|
| {Base Url}/file/{object_name}  | DELETE  | object/file name  |Remove or Delete object from cloud storage|
| {Base Url}/file/upload/        | POST    |Payload bisa di lihat di bawah|Upload or add file to cloud storage|
| {Base Url}/files/page=0&itemPerPage=?| GET | menampilkan semua file yg sudah di upload|

## Payload Upload
Request Body : form/data (postman), jika di hit dari frontend (web,mobile) menggunakan multipart/form-data:
|Key|Value|Description|
|---|----|----|
|upload_type|text|*PHOTO_PROFILE*: upload file image untuk photo_profile, file format yg didukung adalah png dan jpg <br/> *UPLOAD_CONTENT*: upload image untuk content <br/> *DOCUMENT*: upload file document pdf dan docx\|
|file|image.png or photo.png| masukan file yg akan di upload|
