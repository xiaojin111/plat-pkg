package microconfig

import (
	"crypto/rsa"

	"github.com/jinmukeji/plat-pkg/jwt/keystore"
	"github.com/micro/go-micro/config"

	"github.com/jinmukeji/plat-pkg/jwt"
)

type keyItem struct {
	keyPath   []string
	publicKey *rsa.PublicKey

	CfgID          string `json:"id" yaml:"id"`
	CfgFingerprint string `json:"fingerprint" yaml:"fingerprint"`
	CfgPublicKey   string `json:"public_key" yaml:"public_key"`
	Disabled       bool   `json:"disabled" yaml:"disabled"`
}

var _ keystore.KeyItem = (*keyItem)(nil)

func (i keyItem) ID() string {
	return i.CfgID
}

func (i keyItem) PublicKey() *rsa.PublicKey {
	return i.publicKey
}

func (i keyItem) Fingerprint() string {
	return i.CfgFingerprint
}

type MicroConfigStore struct {
	baseCfg    []string
	cachedKeys map[string]*keyItem
}

var _ keystore.Store = (*MicroConfigStore)(nil)

func NewMicroConfigStore(baseCfg ...string) *MicroConfigStore {
	return &MicroConfigStore{
		baseCfg:    baseCfg,
		cachedKeys: make(map[string]*keyItem),
	}
}

func (s *MicroConfigStore) Get(id string) keystore.KeyItem {
	// 首先从缓存中找
	if k, ok := s.cachedKeys[id]; ok && !k.Disabled {
		return k
	}

	item, err := s.loadFromConfig(id)
	if err != nil {
		return nil
	}

	// 添加到缓存
	s.cachedKeys[id] = item

	if item.Disabled {
		return nil
	}

	return item
}

func (s *MicroConfigStore) loadFromConfig(id string) (*keyItem, error) {
	item := keyItem{}

	// 从配置中找
	path := s.keyPath(id)
	if err := config.Get(path...).Scan(&item); err != nil {
		return nil, err
	}

	// found
	item.keyPath = path
	publicKey, err := jwt.LoadRSAPublicKey([]byte(item.CfgPublicKey))
	if err != nil {
		return nil, err
	}
	item.publicKey = publicKey

	return &item, nil
}

func (s *MicroConfigStore) keyPath(id string) []string {
	// <base_key_path>/:id
	keyPath := append(s.baseCfg, id)
	return keyPath
}
