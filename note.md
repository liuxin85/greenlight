postgres://greenlight:pa55word@localhost/greenlight

### windows golang-migration

1. Download pre-build binary from [https://github.com/golang-migrate/migrate/releases]

2. Unzip the file and put migrate.exe to C:\Program Files\Go\bin

3. migrate -version

migrate command different in windows from linux.
In windows:
migrate create -seq -ext sql -dir migrations create_movies_table  

In Linux:
migrate create -seq -ext=.sql -dir=./migrations create_movies_table


$env:GREENLIGHT_DB_DSN="postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable"
echo $env:GREENLIGHT_DB_DSN


psql $env:GREENLIGHT_DB_DSN

DROP TABLE IF EXISTS schema_migrations;
DROP TABLE IF EXISTS movies;

migrate -path migrations -database "$env:GREENLIGHT_DB_DSN" up 


# In git bash
migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
source .env
migrate -path ./migrations -database $GREENLIGHT_DB_DSN up


# Global Rate Limiting
go get golang.org/x/time/rate@latest