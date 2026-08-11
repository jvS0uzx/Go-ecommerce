package handler

import (
	"ecommerce-api/internal/model"
	"ecommerce-api/internal/service"
	
	"net/http"
	"encoding/json"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		http.Error(w, "Dados Inválidos", http.StatusBadRequest)


		return 
	}

	err = h.service.CreateUser(&user)

	if err != nil{
		http.Error(w, "Erro interno ao cadastrar o usuário", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Usuário criado com sucesso!",
	})
}

