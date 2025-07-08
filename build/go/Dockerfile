# Goの公式イメージを使用
FROM golang:1.24.3-alpine

# 作業ディレクトリを src/ に合わせて設定
WORKDIR /go/app

# パッケージをインストール
RUN apk update && apk add git --no-cache git bash

# src/ の中身をすべてコピー（go.mod が含まれる）
COPY . ./

# go.mod, go.sum を元に依存をダウンロード
RUN go mod download

# アプリケーションをビルド（main.go があることが前提）
RUN go build -o /go/bin/app .

# 環境変数とポート設定
ENV PORT=8080
EXPOSE 8080

# アプリケーション実行
CMD ["/go/bin/app"]
