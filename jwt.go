package main

import (
	"time"
	"fmt"
	"io/ioutil"
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
)

var privkey *rsa.PrivateKey
var pubkey  *rsa.PublicKey

// initKeys: save generated keys in memory
func initKeys() error {
	var err error
	// openssl genrsa -out privkey.pem 2048
	privdata, err := ioutil.ReadFile("./keys/privkey.pem")
	if err != nil { return err }

	// openssl rsa -in privkey.pem -pubout -out pubkey.pem
	pubdata, err := ioutil.ReadFile("./keys/pubkey.pem")
	if err != nil { return err }

	privkey, err = jwt.ParseRSAPrivateKeyFromPEM(privdata)
	if err != nil { return err }

	// fmt.Println(*privkey)
	
	pubkey, err = jwt.ParseRSAPublicKeyFromPEM(pubdata)
	if err != nil {	return err }

	return nil
}

func calcJWT(uid string, expires time.Time) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": uid,
		// exp = nbf
		// the number of seconds elapsed since January 1, 1970 UTC.
		// int64
		"exp": expires.Unix(),
	})
	
	//var token *jwt.Token = jwt.New(jwt.SigningMethodRS256)
	//token.Claims["uid"] = uid
	//token.Claims["exp"] = expires.Unix()

	return token.SignedString(privkey)
}

func cbkJwtParse(token *jwt.Token) (interface{}, error) {
	
	// Don't forget to validate the alg is what you expect:
	// SigningMethodHMAC or SigningMethodRSA types
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {

		return nil, fmt.Errorf("Unexpected signing method: %v",
			token.Header["alg"])

	} else {
		return pubkey, nil
	}
	//myLookupKey(token.Header["kid"])
}

func checkJWT(inputToken string) (string, error){
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	

	
	//*jwt.Token
	token, err := jwt.Parse(inputToken, cbkJwtParse)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["foo"], claims["nbf"])
		return claims["uid"].(string), nil
	} else {
		return "", err
	}
}
