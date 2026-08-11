package repository

import (
	"database/sql"
	"log"

	"ecommerce-api/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {

	query := `SELECT article_id, product_code, prod_name, product_type_no, product_type_name, image_url, created_at FROM product LIMIT 50`


		rows, err := r.db.Query(query)
		if err != nil {
			log.Println("Erro ao buscar produtos: ", err)
			return nil, err
		}
		defer rows.Close()

		var products []model.Product

		for rows.Next() {
			var p model.Product

			err := rows.Scan(
				&p.ArticleID, 
				&p.ProductCode, 
				&p.ProdName, 
				&p.ProductTypeNo, 
				&p.ProductTypeName, 
				&p.ImageURL, 
				&p.CreatedAt,
			)

			if err != nil {
				log.Println("Erro ao ler linha: ", err)
				continue
			}

			products = append(products, p)

	}

	return products, nil

}