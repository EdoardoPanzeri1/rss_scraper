package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EdoardoPanzeri1/rss_scraper/internal/database"
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
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handlerFeedsGet)

	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	mux.HandleFunc("GET /v1/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

	mux.HandleFunc("GET /v1/healtz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	const collectionCurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionCurrency, collectionInterval)

	log.Printf("Serving file on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
