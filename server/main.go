package main

import (
	"fmt"
	"net/http"
	"server/helper"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Start main")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Printf("Welcome home")
	})
	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
