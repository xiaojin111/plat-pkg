package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// LoadRSAPrivateKeyFromPEM 加载 RSA 私钥
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

// LoadRSAPublicKeyFromPEM 加载 RSA 公钥
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

// CreateClaims 根据 APP ID 与过期时间间隔创建一个 JWT Claims
func CreateClaims(appId string, inr time.Duration) *jwt.StandardClaims {
	// Create the Claims
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),          // iat
		ExpiresAt: now.Add(inr).Unix(), // exp
		Issuer:    appId,               // iss
	}

	return claims
}

// RS256SignJWT 使用 RS256 算法对 claims 进行签名
func RS256SignJWT(claims *jwt.StandardClaims, key *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	return ss, err
}

// RSAVerifyJWT 使用 RSA 算法（RS256/RS384/RS512) 对 JWT Token 进行验证.
// maxExpInterval 为最大过期时间间隔，单位为秒.
func RSAVerifyJWT(tokenString string, key *rsa.PublicKey, maxExpInterval int64) (bool, error) {

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
