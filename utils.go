package mpesa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
  "fmt"
	"io/ioutil"
)

// GetSecurityCredential generates a security credential
// The security credential depends on the environment set
func (s Service) GetSecurityCredential(initiatorPassword string) (string, error) {
	var pubKey []byte

	var fileName string
	if s.Env == PRODUCTION {
		fileName = "certs/live.cer"
	} else {
		fileName = "certs/sandbox.cer"
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

// GeneratePassword by base64 encoding BusinessShortcode, Passkey, and Timestamp
func GeneratePassword(shortCode, passkey, timestamp string) string {
	str := fmt.Sprintf("%s%s%s", shortCode, passkey, timestamp)
	return base64.StdEncoding.EncodeToString([]byte(str))
}
