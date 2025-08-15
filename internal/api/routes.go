package api

import (
	"github.com/baobei23/lk21-go/internal/api/handler"
	"github.com/baobei23/lk21-go/internal/api/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.CacheMiddleware)

	router.HandleFunc("/movies", handler.LatestMovies).Methods("GET")
	router.HandleFunc("/popular/movies", handler.PopularMovies).Methods("GET")
	router.HandleFunc("/recent-release/movies", handler.RecentReleaseMovies).Methods("GET")
	router.HandleFunc("/top-rated/movies", handler.TopRatedMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", handler.MovieDetails).Methods("GET")

	router.HandleFunc("/series", handler.LatestSeries).Methods("GET")
	router.HandleFunc("/popular/series", handler.PopularSeries).Methods("GET")
	router.HandleFunc("/recent-release/series", handler.RecentReleaseSeries).Methods("GET")
	router.HandleFunc("/top-rated/series", handler.TopRatedSeries).Methods("GET")
	router.HandleFunc("/series/{id}", handler.SeriesDetails).Methods("GET")

	return router
}
