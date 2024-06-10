FROM golang:1.19.3

WORKDIR /app/theater-api

# MySQLクライアントのインストール
RUN apt-get update && apt-get install -y default-mysql-client

# 開発ツールのインストール
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download


RUN go install github.com/cosmtrek/air@v1.40.4

# サービスの開始コマンド
CMD ["make", "start"]
