package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAService interface {
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
	GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey)
	Write() error
	Read() error
}

type RSAServiceImpl struct {
	privatePath string
	publicPath  string
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}

var _ RSAService = &RSAServiceImpl{}

func NewRSAService(privPath string, pubPath string) *RSAServiceImpl {
	return &RSAServiceImpl{
		privatePath: privPath,
		publicPath:  pubPath,
	}
}

func (r *RSAServiceImpl) GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil
	}
	r.privateKey = privkey
	r.publicKey = &privkey.PublicKey
	return privkey, &privkey.PublicKey
}

func (r *RSAServiceImpl) GetPrivateKey() *rsa.PrivateKey {
	return r.privateKey
}

func (r *RSAServiceImpl) GetPublicKey() *rsa.PublicKey {
	return r.publicKey
}

func (r *RSAServiceImpl) Write() error {
	privPem := exportRsaPrivateKeyAsPemStr(r.privateKey)
	pubPem, err := exportRsaPublicKeyAsPemStr(r.publicKey)
	if err != nil {
		return fmt.Errorf("exporting public key as pem string: %w", err)
	}

	err = os.WriteFile(r.privatePath, []byte(privPem), 0600)
	if err != nil {
		return fmt.Errorf("writing private key to file: %w", err)
	}

	err = os.WriteFile(r.publicPath, []byte(pubPem), 0600)
	if err != nil {
		return fmt.Errorf("writing public key to file: %w", err)
	}

	return nil
}

func (r *RSAServiceImpl) Read() error {
	privKey, err := readPemPrivFile(r.privatePath)
	if err != nil {
		return fmt.Errorf("reading private key from file: %w", err)
	}
	r.privateKey = privKey

	pubKey, err := readPemPubFile(r.publicPath)
	if err != nil {
		return fmt.Errorf("reading public key from file: %w", err)
	}
	r.publicKey = pubKey

	return nil
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)
	return string(privkeyPem)
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPem), nil
}

func readPemPubFile(path string) (*rsa.PublicKey, error) {
	// Read the key file
	pubPEM, err := os.ReadFile(path) // #nosec G304 -- file path is controlled
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func readPemPrivFile(path string) (*rsa.PrivateKey, error) {
	privPEM, err := os.ReadFile(path) // #nosec G304 -- file path is controlled
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}
