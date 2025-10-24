package main

import (
	"database/sql"
	"fullstack/movie/data"
	"fullstack/movie/handlers"
	"fullstack/movie/logger"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	//env load
	if err := godotenv.Load(); err != nil {
		logInstance.Error("Error loading .env file", err)
		log.Fatal("No .env file found")
	}

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		logInstance.Error("DATABASE_URL not set in environment variables", nil)
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		logInstance.Error("Error opening database connection", err)
		log.Fatal("Error opening database connection")
	}
	logInstance.Info("Database connection string loaded")
	defer db.Close()

	// Initialize repositories
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		logInstance.Error("Failed to initialize movie repository", err)
		log.Fatal("Failed to initialize movie repository")
	}

	//Backend APIs
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	// Static files or Frontend
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Start server
	const addr = ":8080"
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	logInstance.Error("Hello from Error system", nil)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}
