package jwt

import (
	"context"
	"crypto/rsa"
	"errors"
	"strings"

	mlog "gitee.com/jt-heath/go-pkg/v2/log"
	"gitee.com/jt-heath/plat-pkg/v2/auth/jwt"
	"gitee.com/jt-heath/plat-pkg/v2/auth/jwt/keystore"
	mc "gitee.com/jt-heath/plat-pkg/v2/auth/jwt/keystore/micro-config"
	"gitee.com/jt-heath/plat-pkg/v2/micro/meta"

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

			token := meta.JwtFromContext(ctx)
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

			ctx = meta.ContextWithAppId(ctx, claims.Issuer)

			err = fn(ctx, req, rsp)
			return err
		}
	}
}
