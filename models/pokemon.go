package models

var (
	// Pokemons ...
	Pokemons map[string]*Pokemon
)

// Pokemon ...
type Pokemon struct {
	ID    string
	Name  string
	Image string
}

func init() {
	Pokemons = make(map[string]*Pokemon)
	Pokemons["one"] = &Pokemon{"id_one", "pokemon one", "image one"}
	Pokemons["two"] = &Pokemon{"id_two", "pokemon two", "image two"}
}

// GetAllPokemon ...
func GetAllPokemon() map[string]*Pokemon {
	return Pokemons
}
