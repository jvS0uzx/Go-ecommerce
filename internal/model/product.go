package model

import "time"

// Product representa o produto o Ecommerce 
type Product struct {
    ArticleID       string    `json:"article_id"`
    ProductCode     string    `json:"product_code"`
    ProdName        string    `json:"prod_name"`
    ProductTypeNo   int       `json:"product_type_no"`
    ProductTypeName string    `json:"product_type_name"`
    ImageURL        string    `json:"image_url"`
    CreatedAt       time.Time `json:"created_at"`
}