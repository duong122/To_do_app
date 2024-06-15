package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/controler"
	"server/helper"
	"server/repository"
	"server/router"
	"server/service"
)

func main() {
	fmt.Printf("Start main")
	//database
	db := config.DatabaseConnection()

	// repository
	todoRepository := repository.NewTodoRepository(db)

	//service
	todoService := service.NewTodoServiceImpl(todoRepository)

	// controller
	todoController := controler.NewTodoController(todoService)

	// router
	routes := router.NewRouter(todoController)
	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
