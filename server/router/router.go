package router

import (
	"fmt"
	"net/http"
	"server/controler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoController *controler.TodoControler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Printf("Welcome home")
	})

	router.GET("/api/todo", todoController.FindAll)
	router.GET("/api/todo/:todoId", todoController.FindById)
	router.POST("/api/todo", todoController.Create)
	router.PATCH("/api/book/:todoId", todoController.Update)
	router.DELETE("/api/todo/:todoId", todoController.Delete)

	return router
}
