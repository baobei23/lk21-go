package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/baobei23/lk21-go/internal/scraper"
	"github.com/gorilla/mux"
)

func SearchedMoviesOrSeries(w http.ResponseWriter, r *http.Request) {

	title := mux.Vars(r)["title"]

	url := fmt.Sprintf("%s/?s=%s", os.Getenv("LK21_URL"), title)

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

	searched, err := scraper.ScrapeSearchedMoviesOrSeries(r, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searched)

}
