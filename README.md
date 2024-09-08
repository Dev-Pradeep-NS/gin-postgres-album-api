
# Gin PostgreSQL Album API
This project is a simple RESTful API built with Go, using the Gin framework and PostgreSQL as the database. The API allows you to manage album information, including retrieving, adding, and updating album records.
## Features
- List all albums
- Retrieve an album by ID
- Add a new album
- Update an existing album
- Delete an album
## Configure PostgreSQL
### Set up a PostgreSQL database and user:
1. #### Open the PostgreSQL terminal:
```bash
sudo -u postgres psql
```
2. #### Create a new database and user:
```sql
CREATE DATABASE albumdb;
CREATE USER albumuser WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE albumdb TO albumuser;
```
3. #### Exit the PostgreSQL terminal:
```sql
\q
```
### Update the database connection details in main.go if necessary:
```go
dsn := "host=localhost user=albumuser password=password dbname=albumdb port=5432 sslmode=disable TimeZone=Asia/Calcutta"
```
## Run Application
Start the server with:
```bash
go run main.go
```
The server will run on localhost:8000.
## API Endpoints
### Retrieve a list of all albums.
```http
GET /albums
```
#### Example Request
```
curl http://localhost:8000/albums
```
### Retrieve a single album by its ID.
#### Request
```http
GET /albums/:id
```
#### Example Request
```bash
curl http://localhost:8000/albums/1
```
### Add a new album.
#### Request
```http
POST /albums
```
#### Example Request
```bash
curl -X POST http://localhost:8000/albums \
-H "Content-Type: application/json" \
-d '{
  "title": "New Album",
  "artist": "New Artist",
  "price": 29.99
}'
```
### Update an existing album.
#### Request
```http
PUT /albums/:id
```
#### Example Request
```bash
curl -X PUT http://localhost:8000/albums/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Album",
  "artist": "Updated Artist",
  "price": 39.99
}'
```
### Delete an album
#### Request
```http
DELETE /albums/:id
```
#### Example Request
```bash
curl -X DELETE http://localhost:8000/albums/1
```