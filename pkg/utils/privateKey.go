package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParsePrivateKey(privateKeyBytes []byte, passphrase []byte) (*rsa.PrivateKey, error) {
	privateBlock, _ := pem.Decode(privateKeyBytes)
	if privateBlock == nil || privateBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	var decryptedBlock []byte
	var err error

	if x509.IsEncryptedPEMBlock(privateBlock) {
		decryptedBlock, err = x509.DecryptPEMBlock(privateBlock, passphrase)
		if err != nil {
			return nil, fmt.Errorf("Error decrypting PEM block: %v\n", err)
		}
	} else {
		decryptedBlock = privateBlock.Bytes
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	return privateKey, nil
}
