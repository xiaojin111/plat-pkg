package keystore

import (
	"crypto/rsa"
	"io/ioutil"
	"path"
	"strings"

	"github.com/jinmukeji/plat-pkg/jwt"
)

type fileKeyItem struct {
	id string

	fpFile string
	fp     string

	publicKey     *rsa.PublicKey
	publicKeyFile string
}

var _ KeyItem = (*fileKeyItem)(nil)

func (i fileKeyItem) ID() string {
	return i.id
}

func (i fileKeyItem) PublicKey() *rsa.PublicKey {
	return i.publicKey
}

func (i fileKeyItem) Fingerprint() string {
	return i.fp
}

type FileStore struct {
	keys map[string]*fileKeyItem
}

var _ Store = (*FileStore)(nil)

func NewFileStore() *FileStore {
	return &FileStore{
		keys: make(map[string]*fileKeyItem),
	}
}

func (s *FileStore) Get(id string) KeyItem {
	if k, ok := s.keys[id]; ok {
		return k
	}
	return nil
}

func (s *FileStore) Load(dir string, id string) error {
	fpFile := path.Join(dir, id+".fp.txt")
	fp, err := loadTextFile(fpFile)
	if err != nil {
		return err
	}
	fp = strings.TrimSpace(fp)

	publicKeyFile := path.Join(dir, id+".pub")
	publicKey, err := jwt.LoadRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		return err
	}

	item := &fileKeyItem{
		id:            id,
		fpFile:        fpFile,
		fp:            fp,
		publicKeyFile: publicKeyFile,
		publicKey:     publicKey,
	}

	s.keys[id] = item

	return nil
}

func loadTextFile(file string) (string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
