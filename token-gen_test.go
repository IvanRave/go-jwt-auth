package main

import (	
	"fmt"
)

func ExampleGenerateToken() {
	tkn := generateToken(5)

	tkn2 := generateToken(5)

	if tkn == tkn2 {
		fmt.Println("not-unique randoms")
	}

	fmt.Println(len(tkn))

	// Output:
	// 5
}
