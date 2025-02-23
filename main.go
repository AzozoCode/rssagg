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

	feed, err := urlToFeed("https://wagslane.dev/index.xml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(feed)

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Error opening database connection. %v", err)
	}

	defer conn.Close()

	db := database.New(conn)

	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

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
	routerV1.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))
	routerV1.Post("/users/feeds", apiCfg.authMiddleware(apiCfg.handlerCreateUserFeed))
	routerV1.Get("/users/feeds", apiCfg.handlerGetUserFeeds)
	routerV1.Post("/users/feeds/follow", apiCfg.authMiddleware(apiCfg.handlerFeedFollowCreate))
	routerV1.Get("/users/feeds/follow", apiCfg.authMiddleware(apiCfg.handlerGetUserFeedFollowById))
	routerV1.Delete("/users/feeds/follow/{feed_id}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeedFollow))
	routerV1.Get("/users/posts", apiCfg.authMiddleware(apiCfg.handlerGetUserPosts))

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
