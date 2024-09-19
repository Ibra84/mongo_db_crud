package main

import (
	"log"
	"mongo_db_crud/internal/db"
	"mongo_db_crud/routes"
	"net/http"
)

func main() {
	session := db.Connect()
	defer session.Close()

	router := routes.NewRouter(session)
	log.Fatal(http.ListenAndServe(":9000", router))
}
