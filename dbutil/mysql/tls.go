package mysql

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
)

func NewTLSConfig(rootCert string) (*tls.Config, error) {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(rootCert)
	if err != nil {
		return nil, err
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, errors.New("Failed to append PEM.")
	}

	return &tls.Config{
		RootCAs: rootCertPool,
	}, nil
}
