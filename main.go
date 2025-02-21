package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/azozocode/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Error opening database connection. %v", err)
	}

	defer db.Close()

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
		ExposedHeaders:   []string{"Link"},
	}))

	routerV1 := chi.NewRouter()
	routerV1.Get("/heathz", handlerReadiness)
	routerV1.Get("/err", handlerErr)
	routerV1.Post("/users", apiCfg.handlerCreateUser)
	routerV1.Get("/users", apiCfg.handlerGetUser)
	router.Mount("/v1", routerV1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server started at %v on port:%v", time.Now(), portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
