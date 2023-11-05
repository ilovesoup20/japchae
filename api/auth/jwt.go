package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("my-secret-key")

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "12345678900"
	claims["name"] = "Charles"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func main() {
	token, err := GenerateToken()
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("Token:", token)

	claims, err := verifyToken(token)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}
	fmt.Println("Claims:", claims)
}
