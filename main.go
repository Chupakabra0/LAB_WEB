package main

import (
	"./models"
	"./routes"
	"./utilities"
	"fmt"
	"net/http"
)

func main() {
	models.InitializeDB()
	utilities.ParseTemplates("templates/*.html")
	http.Handle(routes.UrlPath, routes.NewRouter())
	fmt.Println("Server has been launched...")
	_ = http.ListenAndServe(":8888", nil)
}