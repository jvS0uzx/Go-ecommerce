package main

import (
	"fmt" // É um pacote usado para formartar e imprimir texto no console
	"log"
	"net/http"
	
	
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/service"
	"ecommerce-api/internal/repository"
) 

func main() {
	fmt.Println("Iniciando a API...")


	db := config.ConnectDB()
	defer db.Close()

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	productHand := handler.NewProductHandler(svc)



	r := handler.SetupRoutes(productHand)

	fmt.Println("Server na porta 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal("Erro no servidor: ", err)
	}
}

