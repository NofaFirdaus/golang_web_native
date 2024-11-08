package main

import (
	"net/http"
	"web/config"
	categorycontroller "web/controllers/category_controller"
	homecontroller "web/controllers/home_controller"
)

func main() {
	config.ConnectDb()
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	http.ListenAndServe(":8000", nil)
}
