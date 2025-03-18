package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const _baseURL = "https://pokeapi.co/api/v2/"

var ErrPokemonNotFund = errors.New("Pokemon not found")

type Service struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) (*Service, error) {
	if httpClient == nil {
		return nil, errors.New("empty http client")
	}

	return &Service{httpClient}, nil
}

func (s *Service) GetByID(ctx context.Context, pokemonID string) (*Info, error) {
	url := fmt.Sprintf("%s%s%s/", _baseURL, "pokemon/", pokemonID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "generating http request")
	}

	httpResponse, err := s.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "doing http call")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode == http.StatusNotFound {
		return nil, ErrPokemonNotFund
	}

	rawResponseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading http response body")
	}

	var response pokemonInfoResponse
	if err := json.Unmarshal(rawResponseBody, &response); err != nil {
		return nil, errors.Wrap(err, "unmarshaling response body")
	}

	pokemonInfo := Info{
		ID:   response.ID,
		Name: response.Name,
	}

	return &pokemonInfo, nil
}
