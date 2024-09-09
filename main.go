package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"
)

func main() {
	fmt.Println("started go...")
	// database
	db := config.DatabaseConnection()
	// repository
	bookRepository := repository.NewBookRepository(db)
	// service
	bookService := service.NewBookServiceImple(bookRepository)
	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)
	server := http.Server{Addr: "0.0.0.0:3000", Handler: routes}

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}
