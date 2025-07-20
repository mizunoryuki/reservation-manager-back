package main

import (
	"log"
	"net/http"
	"os"

	"reservation-manager/db"
	"reservation-manager/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
        log.Println("Warning: .env file not found")
    }

    log.Printf("DB_HOST=%s, DB_PORT=%s\n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	client := db.NewClient()

	mux:= routes.InitRoutes(client.Q)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // フロントエンドのURL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	log.Println("Listening on :4000")
	http.ListenAndServe(":4000", handler)
}
