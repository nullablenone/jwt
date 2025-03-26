package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidasiToken(tokenString string) (*jwt.Token, error) {
	// Pastikan token tidak kosong
	if tokenString == "" {
		return nil, errors.New("token tidak boleh kosong")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma yang digunakan sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode tanda tangan tidak valid")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Pastikan token valid
	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	// Ambil claims (data dari token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("gagal mengambil claims dari token")
	}

	// Cek apakah token sudah expired
	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return nil, errors.New("token sudah expired")
		}
	}

	// Jika lolos semua validasi, return token
	return token, nil
}
