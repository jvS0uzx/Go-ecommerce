package service

import(
	"ecommerce-api/internal/model"
	"ecommerce-api/internal/repository"
)

type ProductService struct {
	// Em struct ao declarar variaveis não usamos "=" como variaveis comuns, a varivel + 'espaço' + tipo
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll() ([]model.Product, error){
	return s.repo.GetAll()
}