package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xafm/go-crud/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
