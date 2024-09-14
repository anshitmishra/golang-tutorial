package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User struct represents the model for our user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func init() {
	// Open a connection to the MySQL database
	connectionString := "root:Anshit123@#@tcp(127.0.0.1:3306)/crud"
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Execute the query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Unable to scan rows", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Convert the users slice to JSON
	json.NewEncoder(w).Encode(users)
}

// CreateUser adds a new user to the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Execute the query
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to execute query", http.StatusInternalServerError)
		return
	}

	// Get the ID of the inserted user
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to get LastInsertId", http.StatusInternalServerError)
		return
	}

	// Return the ID of the newly created user
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// UpdateUser updates an existing user in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userID := params["id"]

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Execute the query
	_, err = db.Exec("UPDATE users SET name=?, age=? WHERE id=?", user.Name, user.Age, userID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to execute query", http.StatusInternalServerError)
		return
	}

	// Return success message
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User updated successfully"})
}

// DeleteUser deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userID := params["id"]

	// Execute the query
	_, err := db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to execute query", http.StatusInternalServerError)
		return
	}

	// Return success message
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User deleted successfully"})
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", DeleteUser).Methods("DELETE")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
