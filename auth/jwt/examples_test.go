package jwt_test

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/jinmukeji/plat-pkg/v2/auth/jwt"
)

func ExampleRSAVerifyCustomJWT() {
	// MyClaims is a custom claims
	type MyClaims struct {
		jwt.StandardClaims

		AccessToken string `json:"access_token"`
	}

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQwNTMyMTQsImlhdCI6MTU5NDA1MjYxNCwiaXNzIjoiYXBwLXRlc3QxIn0.Xj2bALCrcIMHLHmeeI7ipRddoxU21MmigH3EBr9T_wygkZiZyzOOs-KU2VKuwMhnVsI0vU1iQKs0lCoHt8hSUGddHBjQ4oXcgfo9LWeKl0mluAeVzuBVsI-cZqDAapn5vKRrHvw2IsF-luJNB9th9-HY3_4Nif7OOKGc7DoYkzy-gazKl1lqOH76cy9jQBZ_FNYyKKh28_FgBECxoOogAfakyclPLfXjIxqvpAMMYYp3x0Gbeb1NtRToLNEHeJBEAs1W3vgCQ9i3DF2F1PP3XKHWifUp6MANMgt3w1ghPxxUK2MRHe1oX6wnu652GtspKQ0EJq5GnWMTie0KdRZCfw"
	key, err := jwt.LoadRSAPublicKeyFromPEM("public_key.pem")
	if err != nil {
		panic(err)
	}

	opt := jwt.VerifyOption{
		MaxExpInterval: 10 * time.Minute,
		GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
			// ignore iss check

			return key
		},
	}

	claims := MyClaims{}

	valid, err := jwt.RSAVerifyCustomJWT(token, opt, &claims)
	fmt.Printf("IsValid: %v\n", valid)
	if err != nil {
		fmt.Printf("Validation Error: %v\n", err)
	}
	fmt.Println("Claims:", claims)
}
