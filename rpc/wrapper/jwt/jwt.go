package jwt

import (
	"context"
	"crypto/rsa"
	"errors"
	"strings"

	mlog "github.com/jinmukeji/go-pkg/v2/log"
	cm "github.com/jinmukeji/plat-pkg/v2/rpc/ctxmeta"
	"github.com/jinmukeji/plat-pkg/v2/rpc/jwt"
	"github.com/jinmukeji/plat-pkg/v2/rpc/jwt/keystore"
	mc "github.com/jinmukeji/plat-pkg/v2/rpc/jwt/keystore/micro-config"

	"github.com/micro/go-micro/v2/server"
)

var (
	log *mlog.Logger = mlog.StandardLogger()
)

type jwtWrapper struct {
	opt   Options
	store keystore.Store
}

func newJwtWrapper(opt Options) *jwtWrapper {
	w := jwtWrapper{
		opt: opt,
	}

	baseKeyPath := splitPath(opt.MicroConfigPath)

	store := mc.NewMicroConfigStore(baseKeyPath...)
	w.store = store

	return &w
}

func splitPath(p string) []string {
	s := strings.Trim(p, "/")
	return strings.Split(s, "/")
}

var (
	ErrInvalidJWT = errors.New("forbidden: JWT is invalid")
)

func NewHandlerWrapper(opt Options) server.HandlerWrapper {
	w := newJwtWrapper(opt)

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {

			token := cm.JwtFromContext(ctx)
			log.Debugf("Received JWT Token: %s", token)

			opt := jwt.VerifyOption{
				MaxExpInterval: w.opt.MaxExpInterval,
				GetPublicKeyFunc: func(iss string) *rsa.PublicKey {
					log.Debugf("Issuer from JWT: %s", iss)
					if key := w.store.Get(iss); key != nil {
						return key.PublicKey()
					}
					return nil
				},
			}

			valid, claims, err := jwt.RSAVerifyJWT(token, opt)
			if !valid {
				log.Warnf("failed to validate JWT: %v", err)
				return ErrInvalidJWT
			}

			ctx = cm.ContextWithAppId(ctx, claims.Issuer)

			err = fn(ctx, req, rsp)
			return err
		}
	}
}
