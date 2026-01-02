package services_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"
)

type RSAServiceTestSuite struct {
	suite.Suite
}

func TestRSAService(t *testing.T) {
	rsaServiceSuite := new(RSAServiceTestSuite)
	suite.Run(t, rsaServiceSuite)
}

func (s *RSAServiceTestSuite) TestGenerateKeyPair() {
	tp, _ := filepath.Abs("../../../test/")
	publicKeyPath := tp + "/key.pem.pub"
	privateKeyPath := tp + "/key.pem"
	rs := services.NewRSAService(privateKeyPath, publicKeyPath)

	privKey, pubKey := rs.GenerateKeyPair()

	s.NotNil(privKey)
	s.NotNil(pubKey)
}

func (s *RSAServiceTestSuite) TestWrite() {
	tp, _ := filepath.Abs("../../../test/")
	publicKeyPath := tp + "/key.pem.pub"
	privateKeyPath := tp + "/key.pem"
	rs := services.NewRSAService(privateKeyPath, publicKeyPath)

	rs.GenerateKeyPair()
	err := rs.Write()

	s.NoError(err)
	s.FileExists(privateKeyPath)
	s.FileExists(publicKeyPath)
	os.Remove(privateKeyPath)
	os.Remove(publicKeyPath)
}

func (s *RSAServiceTestSuite) TestRead() {
	tp, _ := filepath.Abs("../../../test/")
	publicKeyPath := tp + "/key.pem.pub"
	privateKeyPath := tp + "/key.pem"
	rs := services.NewRSAService(privateKeyPath, publicKeyPath)

	privKey, pubKey := rs.GenerateKeyPair()
	rs.Write()
	err := rs.Read()
	privKeyReaded := rs.GetPrivateKey()
	pubKeyReaded := rs.GetPublicKey()

	s.NoError(err)
	s.Equal(privKey.D, privKeyReaded.D)
	s.Equal(pubKey.N, pubKeyReaded.N)
	os.Remove(privateKeyPath)
	os.Remove(publicKeyPath)
}
