package controllers

import (
	"encoding/json"
	"fmt"
	"mongo_db_crud/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves a user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Failed to serialize user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	u.Id = bson.NewObjectId()

	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// UpdateUser updates a user by ID
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	if err := uc.session.DB("mongo-golang").C("users").UpdateId(oid, &u); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s updated\n", id)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user %s\n", id)
}
