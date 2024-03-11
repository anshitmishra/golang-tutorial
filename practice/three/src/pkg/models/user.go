package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/anshitmishra/gosqltodo/src/pkg/config"
)

var db *sql.DB

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	User string `json:"user"`
	QueryStruct
}

type TodoReceive struct {
	Task string `json:"task"`
	User string `json:"user"`
}

type QueryStruct struct {
	CreatedAt sql.NullTime `json:"create_at"`
	UpdatedAt sql.NullTime `json:"update_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllRows() ([]Todo, error) {
	rows, err := db.Query("SELECT * FROM todo WHERE deleted_at IS NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Task, &todo.User, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
			fmt.Println(err)
			return nil, errors.New("empty data")
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetRowById(id int) (Todo, error) {
	var todos Todo
	err := db.QueryRow("SELECT * FROM todo WHERE id = ?", id).
		Scan(&todos.ID, &todos.Task, &todos.User, &todos.CreatedAt, &todos.UpdatedAt, &todos.DeletedAt)
	if err == sql.ErrNoRows {
		return Todo{}, errors.New("User not found")
	} else if err != nil {
		return Todo{}, err
	}
	return todos, nil
}

func CreateRow(data TodoReceive) (string, error) {
	_, err := db.Exec("INSERT INTO `todo`(`task`, `user`, `create_at`, `update_at`, `deleted_at`) VALUES (?,?,?,?,?)", data.Task, data.User, time.Now(), time.Now(), nil)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error")
	}
	return "created", nil
}

func UpdateRowById(data TodoReceive, id int) (string, error) {
	_, err := db.Exec("UPDATE `todo` SET `task`=?,`user`=?,`update_at`=? WHERE id = ?", data.Task, data.User, time.Now(), id)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error")
	}
	return "updated", nil
}

func DeleteRowById(id int) (string, error) {
	_, err := db.Exec("DELETE FROM `todo` WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error")
	}
	return "Deleted", nil
}
