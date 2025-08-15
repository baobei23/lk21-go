package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/baobei23/lk21-go/internal/scraper"
	"github.com/gorilla/mux"
)

func LatestSeries(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/latest-series", os.Getenv("ND_URL"))
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

	series, err := scraper.ScrapeSeries(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

func PopularSeries(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/populer", os.Getenv("ND_URL"))
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

	series, err := scraper.ScrapeSeries(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

func RecentReleaseSeries(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/release", os.Getenv("ND_URL"))
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

	series, err := scraper.ScrapeSeries(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

func TopRatedSeries(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("%s/rating", os.Getenv("ND_URL"))
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

	series, err := scraper.ScrapeSeries(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

func SeriesDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	url := fmt.Sprintf("%s/%s", os.Getenv("ND_URL"), id)

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

	details, err := scraper.ScrapeSeriesDetails(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}
