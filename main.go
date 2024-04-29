package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type media struct {
	Id              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	ReleaseYear     string    `json:"releaseYear"`
	Rating          float32   `json:"rating"`
	numberOfRatings int
	Synopsis        string `json:"synopsis"`
}

func main() {
	media := createMedia()

	var port string
	flag.StringVar(&port, "port", "8000", "port api serves over")
	flag.Parse()

	http.HandleFunc("GET /media", func(w http.ResponseWriter, r *http.Request) {
		marshalledMedia, _ := json.Marshal(media)

		w.Write(marshalledMedia)
	})

	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func createMedia() []media {
	media := []media{
		{
			Title:       "Dune",
			ReleaseYear: "2021",
			Synopsis:    "A noble family becomes embroiled in a war for control over the galaxy's most valuable asset while its heir becomes troubled by visions of a dark future.",
		},
		{
			Title:       "Dune: Part Two",
			ReleaseYear: "2024",
			Synopsis:    "Paul Atreides unites with Chani and the Fremen while seeking revenge against the conspirators who destroyed his family.",
		},
		{
			Title:       "The Shawshank Redemption",
			ReleaseYear: "1994",
			Synopsis:    "Over the course of several years, two convicts form a friendship, seeking consolation and, eventually, redemption through basic compassion.",
		},
	}

	for i := 0; i < len(media); i++ {
		media[i].Id, _ = uuid.NewV7()
	}

	return media
}
