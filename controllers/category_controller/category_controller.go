package categorycontroller

import (
	"html/template"
	"net/http"
	"strconv"
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
		timeString := time.Now().Format("2006-01-02 15:04:05")
		category.CreatedAt = []byte(timeString)
		category.UpdatedAt = []byte(timeString)
		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/add.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		category := categorymodel.Show(id)
		data := map[string]any{
			"categories": category,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		idString, _ := strconv.Atoi(r.FormValue("id"))
		category.Id = uint(idString)
		timeString := time.Now().Format("2006-01-02 15:04:05")
		category.UpdatedAt = []byte(timeString)
		if ok := categorymodel.Update(category); !ok {
			temp, _ := template.ParseFiles("views/category/edit.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		if category := categorymodel.Delete(id); !category {
			http.Redirect(w, r, "/categories", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}
}
