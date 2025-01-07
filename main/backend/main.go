package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"log"
    "github.com/joho/godotenv"
	"os"
)

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("OpenAI API key not set")
    }
	if err := connectDatabase(); err != nil {
        log.Fatalf("Failed to connect to the database: %v\n", err)
    }
    r := mux.NewRouter()

    r.Use(enableCORS)

    r.HandleFunc("/upload", uploadHandler).Methods("POST")
    r.HandleFunc("/images/{sessionID}", getGeneratedImagesHandler).Methods("GET")

    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
