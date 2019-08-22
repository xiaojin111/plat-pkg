package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jinmukeji/plat-pkg/jwt"
)

const (
	// DefaultMaxExpirationInterval 默认最大的过期时间间隔（10分钟）
	DefaultMaxExpirationInterval = 600 // in seconds
)

func main() {
	var keyFile, tokenString string

	flag.StringVar(&keyFile, "key", "", "the RSA public key file with PEM format")
	flag.StringVar(&tokenString, "token", "", "the encoded JWT token")
	flag.Parse()

	key, err := jwt.LoadRSAPublicKeyFromPEM(keyFile)
	die(err)

	valid, err := jwt.RSAVerifyJWT(tokenString, key, DefaultMaxExpirationInterval)
	fmt.Printf("IsValid: %v (%v)\n", valid, err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
