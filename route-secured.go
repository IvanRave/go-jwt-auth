package main

import (
	"fmt"
	"net/http"
)

// all client errors must be pre-defined
var (
	errNotAuthorized = apiError{
		Code: "NotAuthorized",
		Description: "Go to a /code.html page",
	}
)

// Get secured page
func routeSecured(w http.ResponseWriter, r *http.Request) error {
	
	rcookie, err := r.Cookie("authtoken")

	if err != nil {
		return errNotAuthorized
	}	
	
	uid, err := checkJWT(rcookie.Value)

	if err != nil {
		return err
	}

	// todo: verify exp date: delete a cookie if expired
	
	fmt.Fprint(w, uid)
	
	return nil
}
