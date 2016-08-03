package main

import (	
	"fmt"
)

func ExampleGenerateToken() {
	tkn := generateToken(5)

	fmt.Println(len(tkn))

	// Output:
	// 5
}
