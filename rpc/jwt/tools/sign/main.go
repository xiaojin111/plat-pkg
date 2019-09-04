package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jinmukeji/plat-pkg/rpc/jwt"
)

func main() {
	var keyFile, appId, interval string

	flag.StringVar(&keyFile, "key", "", "the RSA private key file with PEM format")
	flag.StringVar(&appId, "iss", "", "the issuer")
	flag.StringVar(&interval, "inr", "120s", "the time interval between iat and exp")
	flag.Parse()

	dur, err := time.ParseDuration(interval)
	die(err)

	key, err := jwt.LoadRSAPrivateKeyFromPEM(keyFile)
	die(err)

	claims := jwt.CreateClaims(appId, dur)

	ss, err := jwt.RS256SignJWT(claims, key)
	die(err)

	fmt.Println(ss)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
