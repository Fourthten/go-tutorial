## Source
```bash
https://glide.readthedocs.io/en/latest/
brew install glide
glide init, Yes

https://www.postgresql.org
brew install postgresql

https://github.com/go-pg/pg
glide get github.com/go-pg/pg
Yes, Minor

https://github.com/pressly/goose
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## PostgreSQL
```bash
To restart postgresql after an upgrade:
  brew services restart postgresql
Or, if you don't want/need a background service you can just run:
  /opt/homebrew/opt/postgresql/bin/postgres -D /opt/homebrew/var/postgres
```

```
psql postgres
\du
CREATE ROLE postgres WITH LOGIN PASSWORD 'postgres';
ALTER ROLE postgres CREATEDB;

createuser postgres --createdb

psql postgres -U postgres
CREATE DATABASE tuts;
GRANT ALL PRIVILEGES ON DATABASE tuts TO postgres;
\list
```

## Getting Started
```bash
go build
./pg_tuts
```

## Migration
```
goose create prod_sec_name_update sql
goose postgres "user=postgres dbname=tuts sslmode=disable password=" status
goose postgres "user=postgres dbname=tuts sslmode=disable password=" up
goose postgres "user=postgres dbname=tuts sslmode=disable password=" down

goose postgres "user=postgres dbname=tuts sslmode=disable password=" up-to 20211208073438
goose postgres "user=postgres dbname=tuts sslmode=disable password=" down-to 0
```

[Database modeling with golang & postgresql](https://www.udemy.com/share/101Kxy3@X5JPGJLhp6QRowThaa97m9BKLFNgpBY2CcBnYs5zrS7ugEOhbCZpaIxJEQJX-V9E/)