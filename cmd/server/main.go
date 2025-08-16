package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/baobei23/lk21-go/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := api.NewRouter()

	// Add the root handler to the router
	router.HandleFunc("/", rootHandler)

	fmt.Printf("[%s] server running\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := map[string]interface{}{
		"message": "Unofficial LK21 (LayarKaca21) and NontonDrama APIs",
		"data": map[string]string{
			"LK21_URL": os.Getenv("LK21_URL"),
			"ND_URL":   os.Getenv("ND_URL"),
		},
	}

	json.NewEncoder(w).Encode(data)
}
