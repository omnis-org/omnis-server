package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

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
