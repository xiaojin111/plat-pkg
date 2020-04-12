package keystore

import (
	"crypto/rsa"
)

type Store interface {
	Get(id string) KeyItem
}

type KeyItem interface {
	// ID 返回 ID 标识
	ID() string

	// PublicKey 返回 RSA 公钥
	PublicKey() *rsa.PublicKey

	// Fingerprint 返回 RSA 秘钥对的指纹
	Fingerprint() string
}
