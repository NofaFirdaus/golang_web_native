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
		if err := rows.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`INSERT INTO products (name,category_id,stock,description,created_at,updated_at) 
	VALUE (?,?,?,?,?,?)`, product.Name, product.CategoryId, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		panic(err)
	}
	last, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return last > 0
}

func Update(product entities.Product) bool {
	result, err := config.DB.Exec(`UPDATE products SET name = ?, category_id = ?,stock = ?,description = ?, updated_at = ? WHERE id = ?`, product.Name, product.CategoryId, product.Stock, product.Description, product.UpdatedAt, product.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0

}

func Show(id int) entities.Product {
	result := config.DB.QueryRow(`SELECT id, name, category_id, stock, description FROM products WHERE id = ?`, id)
	var product entities.Product

	if err := result.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Stock, &product.Description); err != nil {
		panic(err)
	}
	return product
}
