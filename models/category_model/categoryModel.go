package categorymodel

import (
	"web/config"
	"web/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`INSERT INTO categories (name,created_at,updated_at) 
	VALUE (?,?,?)`, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		panic(err)
	}
	last, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return last > 0
}

func Update(category entities.Category) bool {
	result, err := config.DB.Exec(`UPDATE categories SET name = ?, created_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, category.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0

}

func Show(id int) entities.Category {
	result := config.DB.QueryRow(`SELECT id ,name FROM categories WHERE id =? `, id)
	var category entities.Category

	if err := result.Scan(&category.Id, &category.Name); err != nil {
		panic(err)
	}
	return category
}

func Delete(id int) bool {
	result, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}
