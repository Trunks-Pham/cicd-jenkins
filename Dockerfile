FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o myapp .

EXPOSE 3000

# Command to run the application
CMD ["./myapp"]