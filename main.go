package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rss_scraper/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbUrl := os.Getenv("DATABSE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment is not set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)

	mux.HandleFunc("GET /v1/healtz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving file on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
