package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"	
)

var Config Configuration
const defaultConfigFile string = "kiwi.go"

func main() {
	// Requires &Config to be accessible
	err := ImportConf()
	if err != nil {
		fmt.Println("Error Reading configuration: ", err)
	}

	// Initiate the web app
	r := mux.NewRouter()
	//r.HandleFunc("/", App)

	// Serve Static Files via "/"
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(Config.StaticDir + "/")))

	err = http.ListenAndServe(":"+strconv.Itoa(Config.Port), r)
	if err != nil {
		fmt.Println(err)
		return
	}
}