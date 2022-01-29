package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}

func greet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	w.Write([]byte("Hello, " + name))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/greet/:name", greet)
	router.POST("/greet/:name", greet)
	http.ListenAndServe(":8080", router)
}
