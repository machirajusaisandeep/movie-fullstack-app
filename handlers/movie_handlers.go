package handlers

import (
	"encoding/json"
	"fullstack/movie/data"
	"fullstack/movie/logger"
	"net/http"
)

type MovieHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

func NewMovieHandler(logger *logger.Logger, storage data.MovieStorage) *MovieHandler {
	return &MovieHandler{
		logger:  logger,
		storage: storage,
	}
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetTopMovies()
	if err != nil {
		h.logger.Error("Failed to get top movies", err)
		http.Error(w, "Failed to get top movies", http.StatusInternalServerError)
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served top movies")
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetRandomMovies()
	if err != nil {
		h.logger.Error("Failed to get random movies", err)
		http.Error(w, "Failed to get random movies", http.StatusInternalServerError)
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served random movies")
	}
}
