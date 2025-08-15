package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/baobei23/lk21-go/internal/scraper"

	"github.com/gorilla/mux"
)

func LatestMovies(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/latest", os.Getenv("LK21_URL"))
	if page != "1" {
		url = fmt.Sprintf("%s/page/%s", url, page)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	movies, err := scraper.ScrapeMovies(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func PopularMovies(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/populer", os.Getenv("LK21_URL"))
	if page != "1" {
		url = fmt.Sprintf("%s/page/%s", url, page)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	movies, err := scraper.ScrapeMovies(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func RecentReleaseMovies(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/release", os.Getenv("LK21_URL"))
	if page != "1" {
		url = fmt.Sprintf("%s/page/%s", url, page)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	movies, err := scraper.ScrapeMovies(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func TopRatedMovies(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/rating", os.Getenv("LK21_URL"))
	if page != "1" {
		url = fmt.Sprintf("%s/page/%s", url, page)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	movies, err := scraper.ScrapeMovies(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func MovieDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	url := fmt.Sprintf("%s/%s", os.Getenv("LK21_URL"), id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	details, err := scraper.ScrapeMovieDetails(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}
