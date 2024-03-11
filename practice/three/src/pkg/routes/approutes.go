package routes

import (
	"github.com/anshitmishra/gosqltodo/src/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRouters = func(router *mux.Router) {
	router.HandleFunc("/todo", controllers.GetAllTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodoById).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodoById).Methods("DELETE")
}
