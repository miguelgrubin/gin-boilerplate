package services_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"
)

type HashServiceArgon2TestSuite struct {
	suite.Suite
}

func (s *HashServiceArgon2TestSuite) TestHash() {
	hashService := services.NewHashServiceArgon2()

	password := "my_secure_password"
	hashedPassword, err := hashService.Hash(password)

	s.NoError(err)
	s.NotEmpty(hashedPassword)
	s.NotContains(hashedPassword, password)
}

func (s *HashServiceArgon2TestSuite) TestCompareSuccess() {
	hashService := services.NewHashServiceArgon2()

	password := "my_secure_password"
	hashedPassword, err := hashService.Hash(password)
	isEqual := hashService.Compare(hashedPassword, password)

	s.True(isEqual)
	s.NoError(err)
}

func (s *HashServiceArgon2TestSuite) TestCompareFailure() {
	hashService := services.NewHashServiceArgon2()

	password := "my_secure_password"
	wrongPassword := "my_wrong_password"
	hashedPassword, err := hashService.Hash(password)
	isEqual := hashService.Compare(hashedPassword, wrongPassword)

	s.NoError(err)
	s.False(isEqual)
}

func TestHashServiceArgon2(t *testing.T) {
	csts := new(HashServiceArgon2TestSuite)
	suite.Run(t, csts)
}
