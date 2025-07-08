package main

import (
	"log"
	"net/http"
	"os"

	"reservation-manager/db"
	"reservation-manager/migration"
	"reservation-manager/routes"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
        log.Println("Warning: .env file not found")
    }

    log.Printf("DB_HOST=%s, DB_PORT=%s\n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	client := db.NewClient()

	err:= migration.Run(os.Getenv("DB_URL"))
	if err != nil {
		log.Println("Migration faild: %v",err)
	}

	mux:= routes.InitRoutes(client.Q)

	log.Println("Listening on :4000")
	http.ListenAndServe(":4000", mux)
}
