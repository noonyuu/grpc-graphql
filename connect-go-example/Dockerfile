FROM golang:1.22.3

WORKDIR /app/connect-go-example

RUN go install github.com/cosmtrek/air@v1.40.4

EXPOSE 3001
# サービスの開始コマンド
CMD ["make", "start"]