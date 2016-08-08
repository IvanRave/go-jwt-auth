package main

import (
//	"time"
	"net/http"

	"github.com/ivanrave/go-jwt-auth/dbauth"
)

var (
	errLgnRequired = apiError{
		Code: "LgnRequired",
		Description: "Lgn format: min 3 symbols",
	}
	errVcodeRequired = apiError{
		Code: "VcodeRequired",
		Description: "Verification code format: min 3 symbols",
	}
	errVcodeNotExists = apiError{
		Code: "VcodeNotExists",
		Description: "A verification code is not found for this login: probably it expired. Please generate a new verification code.",
	}
	errCodeMismatch = apiError{
		Code: "CodeMismatch",
		Description: "The verification code is not correct",
	}
)

func routeLogin (w http.ResponseWriter, r *http.Request) error {
	lgn := r.FormValue("lgn")
	
	if len(lgn) < 3 { return errLgnRequired }

	vcode := r.FormValue("vcode")

	if len(vcode) < 3 { return errVcodeRequired	}
	
	vcodeNeed, err := dbauth.GetVcode(lgn)

	if err != nil {
		if err == dbauth.ErrLgnNotFound {
			// return err
			return errVcodeNotExists
		}
		
		return err
	}
	
	if len(vcodeNeed) < 3 {	return errVcodeNotExists }

	if vcode != vcodeNeed {
		// add retry
		errRetry := dbauth.AddRetry(lgn)
		if errRetry != nil {
			return errRetry
		}
		return errCodeMismatch
	}

	// // then: Generate JWT token
	// // send to the user as Cookie or smth else
	// expiration := time.Now().Add(24 * time.Hour)
	
	// jwt, err := calcJWT(lgn, expiration)

	// if err != nil {	return err }

	// //time.Now().Add(time.Second * time.Duration(seconds))

	// // *Cookie
	// cookie := http.Cookie{
	// 	Name: "authtoken",
	// 	Value: jwt,
	// 	Expires: expiration,
	// 	//		Secure: true,
	// }
	
	// http.SetCookie(w, &cookie)
	
	// //fmt.Fprintf(w, jwt)
	// http.Redirect(w, r, "./", 302)
	return nil
}
