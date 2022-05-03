package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"tokens-overhead/repository"

	"github.com/golang-jwt/jwt"
)

type jwtTokenImpl struct {
	rsaPub  *rsa.PublicKey
	rsaPriv *rsa.PrivateKey
	keyFunc jwt.Keyfunc
}

func NewJWTTokenImpl(pubKeyPath, privKeyPath string) (repository.TokenRepository, error) {
	rsaPub, err := loadRsaPublicKey(pubKeyPath)
	if err != nil {
		return nil, err
	}

	rsaPriv, err := loadRsaPrivateKey(privKeyPath)
	if err != nil {
		return nil, err
	}

	return &jwtTokenImpl{
		rsaPub:  rsaPub,
		rsaPriv: rsaPriv,
		// keyFunc can map direct for rsaPub, since
		// it is the unique existent key
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			return rsaPub, nil
		},
	}, nil

}

func loadRsaPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening rsa priv key: %s", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	return key, err
}

func loadRsaPublicKey(path string) (*rsa.PublicKey, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening rsa pub key: %s", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	return key, err
}

// https://pkg.go.dev/github.com/golang-jwt/jwt#example-New-Hmac
func (j *jwtTokenImpl) Generate(roles []string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"roles": roles,
		},
	)

	tokenString, err := token.SignedString(j.rsaPriv)
	return tokenString, err
}

func (j *jwtTokenImpl) Validate(token string) (bool, error) {
	_, err := jwt.Parse(token, j.keyFunc)
	if _, ok := err.(*jwt.ValidationError); ok {
		return false, errors.New("wrong signature")
	} else if err != nil {
		return false, err
	}
	return true, nil
}
