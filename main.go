package main

import (
	"log"
	"net/http"
	"web/config"
	categorycontroller "web/controllers/category_controller"
	homecontroller "web/controllers/home_controller"
	productcontroller "web/controllers/product_controller"
)

func main() {
	config.ConnectDb()
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	//product
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	serve := ":8000"
	log.Printf("serve running : http://localhost%s \n", serve)
	http.ListenAndServe(serve, nil)
}
