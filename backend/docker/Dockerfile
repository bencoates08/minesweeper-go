# BACKEND

FROM golang:1.19 as backend

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /minesweeper-go

EXPOSE 8080

CMD [ "/minesweeper-go" ]



# DATABASE

FROM --platform=linux/amd64 mysql:8.0.31 as mysql-db

ENV MYSQL_DATABASE="api-minesweeper-db" \
    MYSQL_ROOT_PASSWORD="api-minesweeper" \
    MYSQL_ROOT_HOST="%"

ADD ./docker/schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
