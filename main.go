package main

import (
	"fmt"
	"praktek-jwt/token"
)

func main() {

	//cek perbedaan

	test_token, err := token.GenerateToken(72846)

	fmt.Println("token pertama : ", test_token)
	fmt.Println("error ", err)

	//setelah di validasi
	token, errorrr := token.ValidasiToken(test_token)
	fmt.Println("token : ", token.Raw)
	fmt.Println("error ", errorrr)

}
