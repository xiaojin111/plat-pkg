package jwt

import (
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	mlog "github.com/jinmukeji/go-pkg/log"
	j "github.com/jinmukeji/plat-pkg/jwt"
	"github.com/jinmukeji/plat-pkg/jwt/keystore"
	fstore "github.com/jinmukeji/plat-pkg/jwt/keystore/file"
)

var (
	log *mlog.Logger
)

func init() {
	log = mlog.StandardLogger()
}

type jwt struct {
	enabled   bool
	headerKey string // HTTP Request Header 中的 jwt 使用的 key
	store     keystore.Store
}

const (
	defaultJwtKey = "x-jwt"
)

func (p *jwt) Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:        "enable_jwt",
			Usage:       "Enable JWT validation",
			EnvVar:      "ENABLE_JWT",
			Destination: &(p.enabled),
		},
		cli.StringFlag{
			Name:        "jwt_key",
			Usage:       "JWT HTTP header key",
			EnvVar:      "JWT_KEY",
			Value:       defaultJwtKey,
			Destination: &(p.headerKey),
		},
	}
}

func (p *jwt) Commands() []cli.Command {
	return nil
}

func (p *jwt) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		if !p.enabled {
			return h
		}

		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			token := r.Header.Get(p.headerKey)
			log.Debugf("Received JWT Token: %s", token)

			opt := j.VerifyOption{
				MaxExpInterval: 600,
				GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
					log.Debugf("Issuer from JWT: %s", iss)
					if key := p.store.Get(iss); key != nil {
						return key.PublicKey()
					}
					return nil
				},
			}

			if valid, err := j.RSAVerifyJWT(token, opt); !valid {
				log.Warnf("failed to validate JWT: %v", err)
				http.Error(rw, fmt.Sprintf("forbidden: %v", err), 403)
				return
			}

			// serve the request
			h.ServeHTTP(rw, r)
		})
	}
}

func (p *jwt) Init(ctx *cli.Context) error {
	return nil
}

func (p *jwt) String() string {
	return "jwt"
}

func NewPlugin() plugin.Plugin {
	return NewJWT()
}

func NewJWT() plugin.Plugin {
	// create plugin
	s := fstore.NewFileStore()
	err := s.Load("../jwt/tools/testdata", "app-test1")
	if err != nil {
		panic(err)
	}

	p := &jwt{
		enabled:   false,
		headerKey: defaultJwtKey,
		store:     s,
	}

	return p
}
