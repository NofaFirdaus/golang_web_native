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
