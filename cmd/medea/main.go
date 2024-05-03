package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rickydodd/media-api/internal/config"

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

	var addr, dbUsername, dbPassword, dbName, dbAddr string
	var port, dbPort int
	flag.StringVar(&addr, "ip", "localhost", "address api serves over")
	flag.IntVar(&port, "port", 8000, "port api serves over")
	flag.StringVar(&dbUsername, "db-username", "postgres", "username to access the database")
	flag.StringVar(&dbPassword, "db-password", "", "password to access the database")
	flag.StringVar(&dbName, "db-name", "postgres", "name of the database")
	flag.StringVar(&dbAddr, "db-ip", "localhost", "address database serves over")
	flag.IntVar(&dbPort, "db-port", 5432, "port database serves over")
	flag.Parse()

	conf, err := config.ConfigBuilder().
		Addr(addr).
		Port(port).
		DbUsername(dbUsername).
		DbPassword(dbPassword).
		DbName(dbName).
		DbAddr(dbAddr).
		DbPort(dbPort).
		Build()
	if err != nil {
		log.Fatalln("error:", err)
	}

	// Connect to PostgreSQL.
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%d/%v?sslmode=disable",
		conf.Db.Username,
		conf.Db.Password,
		conf.Db.Addr,
		conf.Db.Port,
		conf.Db.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("error:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalln("error:", err)
	}

	// Create media table in database, if it does not already exist.
	createMediaTable(db)

	http.HandleFunc("GET /media", func(w http.ResponseWriter, r *http.Request) {
		marshalledMedia, _ := json.Marshal(media)

		w.Write(marshalledMedia)
	})

	serveStr := fmt.Sprintf("%v:%d", conf.Addr, conf.Port)
	err = http.ListenAndServe(serveStr, nil)
	if err != nil {
		log.Fatalln("api:", err)
	}
}

func createMediaTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS media (
    media_id uuid NOT NULL,
    title varchar(255) NOT NULL,
    rating numeric(4, 2) NOT NULL DEFAULT 0,
    rating_count integer NOT NULL DEFAULT 0,
    release_year char(4) NOT NULL,
    synopsis text,
    CONSTRAINT media_pkey PRIMARY KEY (media_id)
  )`

	db.Exec(query)
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
