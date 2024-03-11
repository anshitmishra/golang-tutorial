package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/anshitmishra/gosqltodo/src/pkg/models"
	"github.com/gorilla/mux"
)

var NewTodo models.Todo

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	newTodo, err := models.GetAllRows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	res, _ := json.Marshal(newTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	data, err := models.GetRowById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoReceive
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := models.CreateRow(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoReceive
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	todoID := params["id"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	data, err := models.UpdateRowById(todo, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoid := params["id"]

	id, err := strconv.Atoi(todoid)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := models.DeleteRowById(id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
