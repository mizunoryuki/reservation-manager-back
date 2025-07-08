package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"reservation-manager/db/generated"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	DB *sql.DB
	Q  *generated.Queries
}

func NewClient() *DBClient {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("DB応答なし: %v", err)
	}

	return &DBClient{
		DB: db,
		Q:  generated.New(db),
	}
}
