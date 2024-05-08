This is a toy project for portofolio and exploration.

# How to Run
## 1. Create `.env` file from `.env.example`. Fill atleast the `PGPASSWORD`

## 2. Make dbuild.sh executable
In linux / mac :
```
chmod 755 dbuild.sh
./dbuild.sh
```


For windows, just build the docker image using Dockerfile and then run docker compose. As stated in `dbuild.sh`

## 3. Execute dbuild.sh
by :
```sh
./dbuild.sh
```
or manually
```sh
docker build -t go-market-warehouse-api:v1.0 .

docker compose up
```

## 4. Create User table and fill it with records
Use comment on `main.go`

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL
);


INSERT INTO users (name) VALUES ('harry');
INSERT INTO users (name) VALUES ('jack');
INSERT INTO users (name) VALUES ('winston');
```

## 5. Test if application can be accessed
Go to `locahost` then see if it returns json with same id and name that you insert to database.

# Docker Compose: sometimes the application not wait for postgre

Keep rerun the compose file, I still could not find the right way to wait for postgres service to be started and available.

Things that I tried:
- use `depends_on` and `condition = service_healthy`  which return error `dependency failed to start: container local-pg-16 has no healthcheck configured`

Solution:
- use healthcheck from above, which create new section on the postgres service to define `healthcheck`
