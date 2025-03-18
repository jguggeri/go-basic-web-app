package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/jguggeri/go-basic-web-app/internal/services/pokemon"
)

var _regexpPokemonID = regexp.MustCompile(`^[0-9]{1,4}$`)

type Getter interface {
	GetByID(ctx context.Context, pokemonID string) (*pokemon.Info, error)
}

type PokemonHandler struct {
	pokemonGetter Getter
}

func NewPokemon(g Getter) *PokemonHandler {
	return &PokemonHandler{g}
}

func (p *PokemonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	pokemonID := r.PathValue("id")
	if !_regexpPokemonID.MatchString(pokemonID) {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	pokemonInfo, err := p.pokemonGetter.GetByID(r.Context(), pokemonID)
	if err != nil {
		if errors.Is(err, pokemon.ErrPokemonNotFund) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		log.Fatalln(err.Error()) // TODO: cambiar por package log

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	rawResponse, err := json.Marshal(pokemonInfo)
	if err != nil {
		log.Fatalln(err.Error()) // TODO: cambiar por package log

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(rawResponse)
	w.WriteHeader(http.StatusOK)
}
