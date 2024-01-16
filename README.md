# gRPC Microservices in Go

Implementinig examples from gRPC Microservcies in Go

Repo from the book: https://github.com/huseyinbabal/microservices/tree/main

```
docker run -p 3306:3306 \
 -e MYSQL_ROOT_PASSWORD=verysecretpass \
 -e MYSQL_DATABASE=order mysql
```

DB URL: `root:verysecretpass@tcp(127.0.0.1:3306)/order`

## Run App

```
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order \
APPLICATION_PORT=3000 \
ENV=development \
go run cmd/main.go
```
