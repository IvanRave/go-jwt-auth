package main

import (
	"net/http"
	"github.com/ivanrave/goreth"
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
	
	// validate on db level: login + email or phone number regexp
	// save it in Redis
	err := goreth.SetLoginAndVcode(lgn, vcode, 90)

	if err != nil {
		switch (err){
			// client error
		case goreth.ErrLgnExists:
			return errTokenAlreadyExists
		case goreth.ErrLgnInvalid:
			return apiError{
				Code: "BadRequest",
				Description: goreth.ErrLgnInvalid.Error(),
			}
			// server error
		default:
			return err
		}
	}

	http.Redirect(w, r, "./login.html?vcode=" + vcode, 302)	
	//json.NewEncoder(w).Encode(tkn)
	//fmt.Fprintf(w, tkn)
	return nil
}
