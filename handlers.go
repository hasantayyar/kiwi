package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request) error {
	SetHeaders(w)
	vars := mux.Vars(r)
	fmt.Println(vars)
	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	SetHeaders(w)

	return nil
}

func AppHandler(w http.ResponseWriter, r *http.Request) error {
	SetHeaders(w)
	fmt.Fprint(w, "Asdf")
	return nil
}