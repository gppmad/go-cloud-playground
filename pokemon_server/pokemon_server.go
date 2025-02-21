package pokemonserver

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PokemonServer struct {
	Repository CartRepository
}

func (ps *PokemonServer) ServerHttp(response http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	case "/cart":
		ps.GetCard(response, request)
	}
}

func (ps *PokemonServer) GetCard(response http.ResponseWriter, request *http.Request) {
	// Suppose that we get the cart using query params: mypath/cart?cart_id=123
	cartID, err := strconv.Atoi(request.URL.Query().Get("cart_id"))
	if err != nil {
		http.Error(response, "Invalid cart_id", http.StatusBadRequest)
	}

	cart, exists := ps.Repository.GetCart(cartID)
	// If not exists let's return an empty cart
	if !exists {
		cart = Cart{Items: []CartItem{}}
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(cart)
}
