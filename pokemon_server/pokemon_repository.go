package pokemonserver

type PokemonCard struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CartItem struct {
	Card     PokemonCard `json:"card"`
	Quantity int         `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

// Let's declare an interface so we can define the contract between the repo and web server.
type CartRepository interface {
	GetCart(id int) (Cart, bool)
}

// Concrete implementation
type InMemoryCartRepository struct {
	Carts map[int]Cart
}

func (cr *InMemoryCartRepository) GetCart(id int) (Cart, bool) {
	cart, exists := cr.Carts[id]
	return cart, exists
}
