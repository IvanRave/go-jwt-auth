package main

import (	
	"crypto/rand"
)

// var letterRunes = []rune("0123456789")
// //"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func generateToken(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }

const ltrRandomNumberString = "123456789"

func generateToken(n int) string {
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = ltrRandomNumberString[b % byte(len(ltrRandomNumberString))]
	}
	return string(bytes)
}
