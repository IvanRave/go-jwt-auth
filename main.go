package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	"github.com/ivanrave/go-jwt-auth/dbauth"
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

	err := dbauth.InitPool("127.0.0.1",
		"6379",
		0,
		"")
	if err != nil {
		log.Fatal("Redis: ", err)
	}

	// err = initKeys()
	// if err != nil {
	// 	log.Fatal("Keys: ", err)
	// }	

	// err = checkKey()
	// if err != nil {
	// 	log.Fatal("Redis: ", err)
	// }

	http.Handle("/login", appHandler(routeLogin))
	http.Handle("/code", appHandler(routeCode))
	http.Handle("/", appHandler(routeSecured))
	
	http.Handle("/code.html", http.FileServer(http.Dir("./static")))
	http.Handle("/login.html", http.FileServer(http.Dir("./static")))

	const listenPort string = ":9090"
	fmt.Println("http://127.0.0.1" + listenPort)
	err = http.ListenAndServe(listenPort, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
