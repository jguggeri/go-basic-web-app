package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jguggeri/go-basic-web-app/cmd/server/api/handlers"
	"github.com/jguggeri/go-basic-web-app/internal/services/pokemon"
)

const _port = ":8080"

func main() {
	log.Println("initializing [api] app")

	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}

	pokemonClient, err := pokemon.NewClient(httpClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	pokemonHandler := handlers.NewPokemon(pokemonClient)

	log.Println("dependencies created")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "pong!") }) //nolint:lll
	http.HandleFunc("GET /pokemons/{id}", pokemonHandler.GetByID)

	log.Printf("listening at %s...", _port)

	if err := http.ListenAndServe(_port, nil); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("exiting [api] app")
}
