FROM golang:1.19.3

WORKDIR /app/grpc

RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@v1.40.4

CMD ["make", "start"]
