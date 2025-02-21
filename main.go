package main

import (
	"log"
	"net/http"

	pokemon_server "github.com/gppmad/gocloudplayground/pokemon_server"
)

func NewInMemoryCartRepository() *pokemon_server.InMemoryCartRepository {
	pikachu_card := pokemon_server.PokemonCard{
		ID:    "1",
		Name:  "Pikachu",
		Price: 10.0,
	}
	cart_item := pokemon_server.CartItem{Card: pikachu_card, Quantity: 1}

	cart := pokemon_server.Cart{
		Items: []pokemon_server.CartItem{cart_item},
	}
	carts := make(map[int]pokemon_server.Cart)
	carts[1] = cart

	return &pokemon_server.InMemoryCartRepository{
		Carts: carts,
	}
}

func main() {
	// I need to create the web server that accepts the request.
	// The most critical part for me it's to decouple the web server from the repository.
	// in this way and in the test we can easily plug the real repo implementation.

	memoryRepo := NewInMemoryCartRepository()
	server := pokemon_server.PokemonServer{Repository: memoryRepo}
	handler := http.HandlerFunc(server.ServerHttp)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
