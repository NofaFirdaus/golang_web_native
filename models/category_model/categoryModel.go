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
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdateAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`insert into categories (name,created_at,updated_at) 
	value (?,?,?)`, category.Name, category.CreatedAt, category.UpdateAt)
	if err != nil {
		panic(err)
	}
	last, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return last < 0

}
