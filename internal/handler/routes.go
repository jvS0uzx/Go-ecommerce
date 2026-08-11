package handler

import (
	"encoding/json" // Pacote para lidar com JSON
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(productHand *ProductHandler, userHand *UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		response := map[string]string{
			"message": "pong! API respondendo.",
			"status": "success",
		}

		json.NewEncoder(w).Encode(response)

	})

	r.Get("/produtos", productHand.GetAllProducts)

	r.Post("/users", userHand.CreateUser)

	return r
}