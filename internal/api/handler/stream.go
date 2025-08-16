package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/baobei23/lk21-go/internal/scraper"
	"github.com/gorilla/mux"
)

func StreamMovie(w http.ResponseWriter, r *http.Request) {
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

	stream, err := scraper.ScrapeStreamSources(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stream)
}

func StreamSeries(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	parts := strings.Split(id, "-")
	year := parts[len(parts)-1]
	seriesId := strings.Join(parts[:len(parts)-1], "-")

	query := r.URL.Query()
	season := query.Get("season")
	if season == "" {
		season = "1"
	}
	episode := query.Get("episode")
	if episode == "" {
		episode = "1"
	}

	url := fmt.Sprintf("%s/%s-season-%s-episode-%s-%s", os.Getenv("ND_URL"), seriesId, season, episode, year)

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

	stream, err := scraper.ScrapeStreamSources(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stream)
}
