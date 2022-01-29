package main

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received ", r.URL.Path)
		handler(w, r)
		fmt.Println("Response served ", r.URL.Path)
	}
}

func profiler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println("Time taken : ", elapsed)
		}()
		handler(w, r)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("All the products will be served"))
	case "POST":
		w.Write([]byte("A new project is created"))
	case "DELETE":
		w.Write([]byte("A product is deleted"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	/*
		http.HandleFunc("/", profiler(logger(indexHandler)))
		http.HandleFunc("/products", profiler(logger(productsHandler)))
	*/
	http.HandleFunc("/", chain(indexHandler, logger, profiler))
	http.HandleFunc("/products", chain(productsHandler, logger, profiler))
	http.ListenAndServe(":8080", nil)
}

/*
	GET /products
	GET /products/P101
	POST /products/
	PUT /products/P101
	DELETE /products/P101

*/
