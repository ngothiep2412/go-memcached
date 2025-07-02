# Caching in Go By Using memcached

## Requirements
1. Go
2. Mysql
3. memcached: either a remote instance, local binary or docker container



## Running 
- Run
```
go mod download
```

- Create your .env with DATABASE_URL and MEMCAHCED_URL
- Using docker to run memcached:
```
docker run \
  -d \
  -p {YOUR_PORT}:11211 \
  memcached:1.6.9-alpine
```

Then using:
```
go run main.go
```