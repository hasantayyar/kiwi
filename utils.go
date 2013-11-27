package main

import (
	"net/http"
)

func errorAdapter(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w,r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type","text/html")
	w.Header().Set("X-Powered-By", "Kiwi")
}