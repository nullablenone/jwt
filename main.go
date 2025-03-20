package main

import (
	"fmt"
	"praktek-jwt/token"
)

func main() {
	t, err := token.GenerateToken(1)
	if err != nil {
		fmt.Println(err)
	}

	verif, err := token.ValidasiToken(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(verif)
}
