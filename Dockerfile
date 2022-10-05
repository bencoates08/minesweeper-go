FROM golang:1.19 as build

# Setup folders
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o /minesweeper-go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/minesweeper-go" ]

