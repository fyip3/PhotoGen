package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
	"log"
)

type UploadResponse struct {
    GeneratedImages []GeneratedImage `json:"generated_images"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, "Unable to parse form data", http.StatusBadRequest)
        log.Println("ParseMultipartForm error:", err)
        return
    }

    description := r.FormValue("description")
    if description == "" {
        http.Error(w, "Description is required", http.StatusBadRequest)
        return
    }

    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Could not retrieve file from form-data", http.StatusBadRequest)
        log.Println("FormFile retrieval error:", err)
        return
    }
    defer file.Close()

    images, err := generateImages(description)
    if err != nil {
        http.Error(w, "Failed to generate images", http.StatusInternalServerError)
        log.Println("Image generation error:", err)
        return
    }

    // Store images in the database with a unique session ID
    // sessionID := "unique-session-id" // Replace with actual session ID generation
    // if err := saveGeneratedImages(sessionID, images); err != nil {
    //     http.Error(w, "Failed to save images", http.StatusInternalServerError)
    //     log.Println("Database save error:", err)
    //     return
    // }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(UploadResponse{GeneratedImages: images})
}
func getGeneratedImagesHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sessionID := vars["sessionID"]

    images, err := getImagesBySessionID(sessionID)
    if err != nil {
        http.Error(w, "Images not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(UploadResponse{GeneratedImages: images})
}
