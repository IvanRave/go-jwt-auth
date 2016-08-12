package main

import (
	"time"
	"fmt"
	"log"
)

func ExampleCalcJWT() {
	err := initKeys()
	if err != nil {
		log.Fatal(err)
		return
	}
	expiration := time.Now().Add(1 * time.Hour)
	//.Add(24 * time.Hour)
	authToken, err := calcJWT("qwerty", expiration)

	if err != nil {
		log.Fatal(err)
	}

	uid, err := checkJWT(authToken)

	if err != nil {
		if err.Error() == "Token is expired" {
			log.Fatal("MyError: Token is expired")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(uid)
	
	// Output:
	// qwerty
}
