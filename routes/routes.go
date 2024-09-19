package routes

import (
	"mongo_db_crud/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func NewRouter(session *mgo.Session) *httprouter.Router {
	router := httprouter.New()
	uc := controllers.NewUserController(session)

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.PUT("/user/:id", uc.UpdateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	return router
}
