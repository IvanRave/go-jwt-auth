package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := ah(w, r); err != nil {
		switch e := err.(type) {
		case apiError:
			// set 400 status for client errors
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(e)
		default:
			fmt.Println(err)
			http.Error(w, "ServerError", 500)
		}
	}
}

func main() {

	err := initPool()
	if err != nil {
		log.Fatal("Redis: ", err)
	}

	err = initKeys()
	if err != nil {
		log.Fatal("Keys: ", err)
	}	

	err = checkKey()
	if err != nil {
		log.Fatal("Redis: ", err)
	}

	http.Handle("/login", appHandler(routeLogin))
	http.Handle("/code", appHandler(routeCode))
	http.Handle("/", appHandler(routeSecured))
	
	http.Handle("/code.html", http.FileServer(http.Dir("./static")))
	http.Handle("/login.html", http.FileServer(http.Dir("./static")))
	
	err = http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
