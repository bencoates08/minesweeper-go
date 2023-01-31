# BACKEND

FROM golang:1.19 as backend

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /minesweeper-go

EXPOSE 8080

ENV DB_CONTAINER="minesweeper-database" \
    DB_PASSWORD="minesweeper-database-password" \
    DB_NAME="minesweeper-database-name"

CMD [ "/minesweeper-go" ]



# DATABASE

FROM mysql:8.0.32-oracle as mysql-db

ENV MYSQL_DATABASE="minesweeper-database-name" \
    MYSQL_ROOT_PASSWORD="minesweeper-database-password" \
    MYSQL_ROOT_HOST="%"

ADD ./docker/schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
