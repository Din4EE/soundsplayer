package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	repoTrack "github.com/Din4EE/soundsplayer/internal/repo/track"
	serviceTrack "github.com/Din4EE/soundsplayer/internal/service/track"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	db, err := sql.Open("pgx",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_SSLMODE")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	trackService := serviceTrack.NewService(repoTrack.NewRepository(db))
	mux := http.NewServeMux()
	audioFS := http.FileServer(http.Dir("assets/audio"))
	mux.Handle("/audio/", http.StripPrefix("/audio/", audioFS))
	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/playlist", playlistHandler(trackService))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func playlistHandler(trackService *serviceTrack.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		tracks, err := trackService.GetAll(r.Context())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		playlistTemplate, err := template.ParseFiles("templates/playlist.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = playlistTemplate.Execute(w, tracks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
