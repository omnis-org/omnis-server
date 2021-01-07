package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

// ParsePrivKey allow to parse a private key file
// if the path or the content of the file is invalid return an error
// The function return a pointer of rsa.PrivateKey when valid
func ParsePrivKey(AuthKeyFile string) (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(AuthKeyFile)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile failed <- %v", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("jwt.ParseRSAPrivateKeyFromPEM failed <- %v", err)
	}

	return key, nil
}

// ParsePubKey allow to parse a public key file
// if the path or the content of the file is invalid return an error
// The function return a pointer of rsa.PublicKey when valid
func ParsePubKey(AuthPubFile string) (*rsa.PublicKey, error) {
	pubBytes, err := ioutil.ReadFile(AuthPubFile)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile failed <- %v", err)
	}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("jwt.ParseRSAPublicKeyFromPEM failed <- %v", err)
	}

	return pub, nil
}
