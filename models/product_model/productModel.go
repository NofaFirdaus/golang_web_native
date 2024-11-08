package productmodel

import (
	"web/config"
	"web/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`SELECT * FROM products`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products
}
