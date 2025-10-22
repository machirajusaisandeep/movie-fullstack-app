package main

import (
	"fullstack/movie/handlers"
	"fullstack/movie/logger"
	"log"
	"net/http"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// Static files or Frontend
	http.Handle("/", http.FileServer(http.Dir("public")))

	//Backend APIs
	movieHandler := handlers.MovieHandler{}
	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)

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
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}
