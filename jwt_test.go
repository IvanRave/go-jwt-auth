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
	expiration := time.Now().Add(24 * time.Hour)
	authToken, err := calcJWT("qwerty", expiration)

	if err != nil {
		log.Fatal(err)
		return
	}

	uid, err := checkJWT(authToken)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(uid)
	
	// Output:
	// qwerty
}
