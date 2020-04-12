package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/smallstep/cli/crypto/pemutil"
)

var (
	ErrParseClaimsFailed     = errors.New("failed to parse not standard claims")
	ErrNoPublicKey           = errors.New("no public key to verify JWT")
	ErrEmptyToken            = errors.New("token is empty")
	ErrInvalidPrivateKeyFile = errors.New("invalid private key file")
)

// LoadRSAPrivateKeyFromPEM 加载 RSA 私钥
func LoadRSAPrivateKeyFromPEM(keyFile, passFile string) (*rsa.PrivateKey, error) {
	key, err := pemutil.Read(keyFile, pemutil.WithPasswordFile(passFile))
	if err != nil {
		return nil, err
	}
	if sk, ok := key.(*rsa.PrivateKey); ok {
		return sk, nil
	}

	return nil, ErrInvalidPrivateKeyFile
}

// LoadRASPrivateKey 从字节序列中加载 RSA 私钥
func LoadRASPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM(key)
}

// LoadRSAPublicKeyFromPEM 加载 RSA 公钥
func LoadRSAPublicKeyFromPEM(file string) (*rsa.PublicKey, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadRSAPublicKey(f)
}

// LoadRSAPublicKey 从字节序列中加载 RSA 公钥
func LoadRSAPublicKey(key []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(key)
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
	return signJWT(jwt.SigningMethodRS256, key, claims)
}

// RS384SignJWT 使用 RS84 算法对 claims 进行签名
func RS384SignJWT(claims *jwt.StandardClaims, key *rsa.PrivateKey) (string, error) {
	return signJWT(jwt.SigningMethodRS384, key, claims)
}

// RS512SignJWT 使用 RS512 算法对 claims 进行签名
func RS512SignJWT(claims *jwt.StandardClaims, key *rsa.PrivateKey) (string, error) {
	return signJWT(jwt.SigningMethodRS512, key, claims)
}

func signJWT(method jwt.SigningMethod, key interface{}, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	ss, err := token.SignedString(key)

	return ss, err
}

type GetPublicKeyFunc func(iss string) *rsa.PublicKey

type VerifyOption struct {
	MaxExpInterval   time.Duration    // 最大过期时间间隔，单位为秒.
	GetPublicKeyFunc GetPublicKeyFunc // PublicKey 查找函数
}

// RSAVerifyJWT 使用 RSA 算法（RS256/RS384/RS512) 对 JWT Token 进行验证.
func RSAVerifyJWT(tokenString string, opt VerifyOption) (bool, *jwt.StandardClaims, error) {
	var stdClaims *jwt.StandardClaims

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

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return nil, ErrParseClaimsFailed
		}

		stdClaims = claims

		if opt.GetPublicKeyFunc != nil {
			key := opt.GetPublicKeyFunc(claims.Issuer)

			if key != nil {
				return key, nil
			}
		}

		return nil, ErrNoPublicKey
	})

	if token == nil {
		return false, nil, ErrEmptyToken
	}

	if token.Valid {
		inr := float64(stdClaims.ExpiresAt - stdClaims.IssuedAt)
		if inr > opt.MaxExpInterval.Seconds() {
			return false, nil, fmt.Errorf("expiration interval exceeds the limit: %fs", opt.MaxExpInterval.Seconds())
		}

		return true, stdClaims, nil
	} else {
		return false, nil, err
	}
}
