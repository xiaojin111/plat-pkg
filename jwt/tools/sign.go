package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	var keyFile, appId, interval string

	flag.StringVar(&keyFile, "key", "", "the RSA private key file with PEM format")
	flag.StringVar(&appId, "iss", "", "the issuer")
	flag.StringVar(&interval, "inr", "120s", "the time interval between iat and exp")
	flag.Parse()

	dur, err := time.ParseDuration(interval)
	die(err)

	key, err := LoadRSAPrivateKeyFromPEM(keyFile)
	die(err)

	claims := CreateClaims(appId, dur)

	ss, err := SignJWT(claims, key)
	die(err)

	fmt.Println(ss)
}

func LoadRSAPrivateKeyFromPEM(file string) (*rsa.PrivateKey, error) {
	keyFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func CreateClaims(appId string, duration time.Duration) *jwt.StandardClaims {
	// Create the Claims
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(duration).Unix(),
		Issuer:    "app-test",
	}

	return claims
}

func SignJWT(claims *jwt.StandardClaims, key *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	log.Printf("alg: %s", token.Method.Alg())
	log.Printf("iss: %s", claims.Issuer)
	log.Printf("iat: %d", claims.IssuedAt)
	log.Printf("exp: %d", claims.ExpiresAt)

	return ss, err
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
