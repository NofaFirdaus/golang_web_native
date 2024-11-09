package productcontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	"web/entities"
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
	if r.Method == "POST" {
		var product entities.Product
		product.Name = r.FormValue("name")
		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Stock = uint(stock)
		product.Description = r.FormValue("description")
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}
		product.CategoryId = uint(categoryId)
		timeString := time.Now().Format("2006-01-02 15:04:05")
		product.CreatedAt = []byte(timeString)
		product.UpdatedAt = []byte(timeString)

		if ok := productmodel.Create(product); !ok {
			temp, _ := template.ParseFiles("views/product/add.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		product := productmodel.Show(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}
		if err := temp.Execute(w, data); err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == "POST" {
		var product entities.Product
		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}
		product.Stock = uint(stock)
		product.Description = r.FormValue("description")

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		product.CategoryId = uint(categoryId)
		product.Name = r.FormValue("name")

		idString, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		product.Id = uint(idString)

		timeString := time.Now().Format("2006-01-02 15:04:05")
		product.UpdatedAt = []byte(timeString)

		if ok := productmodel.Update(product); !ok {
			temp, _ := template.ParseFiles("views/product/edit.html")
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		if product := productmodel.Delete(id); !product {
			http.Redirect(w, r, "/product", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/product", http.StatusSeeOther)

	}
}
