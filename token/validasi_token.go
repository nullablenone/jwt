package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ValidasiToken(tokenString string) (*jwt.Token, error) {
	//validasi token tidak boleh kosong
	if tokenString == "" {
		return nil, errors.New("token tidak boleh kosong")
	}

	// Parse token & cek validitasnya
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma yang dipakai sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode tanda tangan tidak valid")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
