package main

import (
	"fmt"
	"praktek-jwt/token"
)

func main() {
	user, err := token.GenerateToken(1)
	if err != nil {
		fmt.Println("gagal generate : ", err)
	}
	fmt.Println("token : ", user)
}
