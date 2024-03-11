package main

import (
	"log"
	"net/http"

	"github.com/anshitmishra/gosqltodo/src/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
