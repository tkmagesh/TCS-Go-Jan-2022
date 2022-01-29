package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.ListenAndServe(":8080", nil)
}

/*
	GET /products
	GET /products/P101
	POST /products/
	PUT /products/P101
	DELETE /products/P101

*/
