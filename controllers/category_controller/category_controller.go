package categorycontroller

import (
	"html/template"
	"net/http"
	"time"
	"web/entities"
	categorymodel "web/models/category_model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdateAt = time.Now()
		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/add.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
