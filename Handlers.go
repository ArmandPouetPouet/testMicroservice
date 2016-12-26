package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"testMicroservice/Data"

	"github.com/gorilla/mux"
)

//Index zpouet
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//UserList Get all users
func UserList(w http.ResponseWriter, r *http.Request) {
	w = BuildJSONResponse(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(Data.GetUsers()); err != nil {
		panic(err)
	}
}

//User Get one user
func User(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	index, err := strconv.Atoi(todoID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(Data.GetUser(index))
}

//UserCreate Create one user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user Data.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w = BuildJSONResponse(w, 422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := Data.CreateUser(user)
	w = BuildJSONResponse(w, http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
