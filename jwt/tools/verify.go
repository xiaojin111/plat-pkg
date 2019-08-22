package main

import (
	"crypto/rsa"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
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

	key, err := LoadRSAPublicKeyFromPEM(keyFile)
	die(err)

	valid, err := VerifyJWT(tokenString, key, DefaultMaxExpirationInterval)
	fmt.Printf("IsValid: %v %v\n", valid, err)
}

func LoadRSAPublicKeyFromPEM(file string) (*rsa.PublicKey, error) {
	keyFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyFile)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func VerifyJWT(tokenString string, key *rsa.PublicKey, maxExpInterval int64) (bool, error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			// Support: RS256, RS384, RS512
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false, errors.New("failed to parse not standard claims")
	}

	log.Printf("alg: %s", token.Method.Alg())
	log.Printf("iss: %s", claims.Issuer)
	log.Printf("iat: %d", claims.IssuedAt)
	log.Printf("exp: %d", claims.ExpiresAt)

	if token.Valid {
		inr := claims.ExpiresAt - claims.IssuedAt
		if inr > maxExpInterval {
			return false, fmt.Errorf("expiration interval exceeds the limit: %ds", maxExpInterval)
		}

		return true, nil
	} else {
		return false, err
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
