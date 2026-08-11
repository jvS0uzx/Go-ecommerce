package handler

import(
	"ecommerce-api/internal/service"
	"net/http"
	"encoding/json"
)


type ProductHandler struct {
	prha *service.ProductService
}

func NewProductHandler(prha *service.ProductService) *ProductHandler {
	return &ProductHandler{prha: prha}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request){
	products, err := h.prha.GetAll()

	if (err != nil){
		http.Error(w, "mensagem", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)


}