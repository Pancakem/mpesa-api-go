package mpesa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
)

// GetSecurityCredential generates a security credential
// The security credential depends on the environment set
func (s Service) GetSecurityCredential(initiatorPassword string) (string, error) {
	var pubKey []byte

	var fileName string
	if s.Env == PRODUCTION {
		fileName = "certs/live.cert"
	} else {
		fileName = "certs/sandbox.cert"
	}

	pubKey, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(pubKey)
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	rng := rand.Reader

	encrypted, err := rsa.EncryptPKCS1v15(rng, rsaPublicKey, []byte(initiatorPassword))
	if err != nil {
		return "", err
	}

	enc := base64.StdEncoding.EncodeToString(encrypted)
	return enc, nil
}
