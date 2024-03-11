package main

import (
	"log"
	"net/http"

	"github.com/anshitmishra/golang-mongodb/src/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id}", uc.DeleteUser)
	http.ListenAndServe(":8000", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		log.Fatal(err)
	}

	return session
}
