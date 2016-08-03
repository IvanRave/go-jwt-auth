package main

import (
	"net/http"
)
	
var (
	errTokenAlreadyExists = apiError{
		Code: "TokenAlreadyExists",
		Description: "Need to wait while current token will be expired",
	}
)

func routeCode(w http.ResponseWriter, r *http.Request) error {
	// r.ParseForm()  // parse arguments		
	lgn := r.FormValue("lgn")
	
	// generate a verification code
	vcode := generateToken(5)
	
	// validate login: email or phone number regexp
	
	// save it in Redis
	saved, err := setLoginAndCode(lgn, vcode)

	if err != nil {
		// server error
		return err
	}

	if saved == false {
		// client error
		return errTokenAlreadyExists
	}

	http.Redirect(w, r, "./login.html?vcode=" + vcode, 302)	
	//json.NewEncoder(w).Encode(tkn)
	//fmt.Fprintf(w, tkn)
	return nil
}
