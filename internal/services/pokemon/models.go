package pokemon

// Public structs.
type (
	Info struct {
		ID   int64
		Name string
	}
)

// Private structs.
type (
	pokemonInfoResponse struct {
		ID   int64  //`json:"id"`
		Name string // `json:"name"`
	}
)
