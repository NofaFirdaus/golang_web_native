package productcontroller

import (
	"html/template"
	"net/http"
	categorymodel "web/models/category_model"
	productmodel "web/models/product_model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	product := productmodel.GetAll()
	data := map[string]any{
		"product": product,
	}
	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}
		temp, err := template.ParseFiles("views/product/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
	// if r.Method == "POST" {
	// 	var category entities.Category
	// 	category.Name = r.FormValue("name")
	// 	timeString := time.Now().Format("2006-01-02 15:04:05")
	// 	category.CreatedAt = []byte(timeString)
	// 	category.UpdatedAt = []byte(timeString)
	// 	if ok := categorymodel.Create(category); !ok {
	// 		temp, _ := template.ParseFiles("views/category/add.html")
	// 		temp.Execute(w, nil)
	// 	}
	// 	http.Redirect(w, r, "/categories", http.StatusSeeOther)
	// }
}
