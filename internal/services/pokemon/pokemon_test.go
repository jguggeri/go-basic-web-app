package pokemon_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jguggeri/go-basic-web-app/internal/services/pokemon"
)

func Test(t *testing.T) {
	s, _ := pokemon.NewClient(&http.Client{})

	i, err := s.GetByID(context.Background(), "1")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(i)
}
