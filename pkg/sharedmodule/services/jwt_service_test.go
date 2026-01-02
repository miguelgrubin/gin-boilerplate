package services_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JWTServiceTestSuite struct {
	suite.Suite
	redisMock  *mocks.MockRedisService
	rsaService services.RSAService
	jwtService services.JWTService
}

func (ts *JWTServiceTestSuite) TestGenerateTokens() {
	ts.redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	token, refresh, err := ts.jwtService.GenerateTokens("user-id", "role")

	ts.NoError(err)
	ts.NotEmpty(token)
	ts.NotEmpty(refresh)
}

func (ts *JWTServiceTestSuite) TestValidateToken() {
	ts.redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	token, _, err := ts.jwtService.GenerateTokens("user-id", "role")
	ts.NoError(err)

	valid := ts.jwtService.ValidateToken(token)
	ts.True(valid)
}

func (ts *JWTServiceTestSuite) TestValidateTokenInvalid() {
	valid := ts.jwtService.ValidateToken("invalid-token")
	ts.False(valid)
}

func (ts *JWTServiceTestSuite) TestDecodeToken() {
	ts.redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	token, _, err := ts.jwtService.GenerateTokens("user-id", "role")
	ts.NoError(err)

	data, err := ts.jwtService.DecodeToken(token)
	ts.NoError(err)
	ts.Equal("user-id", data.UserID)
	ts.Equal("role", data.Role)
}

func (ts *JWTServiceTestSuite) TestRefreshToken() {
	ts.redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ts.redisMock.On("Has", mock.Anything).Return(true, nil)
	ts.redisMock.On("Del", mock.Anything).Return(nil)

	_, refresh, err := ts.jwtService.GenerateTokens("user-id", "role")
	ts.NoError(err)

	newToken, newRefresh, err := ts.jwtService.RefreshToken(refresh, "user-id", "role")

	ts.NoError(err)
	ts.NotEmpty(newToken)
	ts.NotEmpty(newRefresh)
}

func (ts *JWTServiceTestSuite) TestInvalidateToken() {
	ts.redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ts.redisMock.On("Has", mock.Anything).Return(true, nil)
	ts.redisMock.On("Del", mock.Anything).Return(nil)

	token, _, err := ts.jwtService.GenerateTokens("user-id", "role")
	ts.NoError(err)

	err = ts.jwtService.InvalidateToken(token)
	ts.NoError(err)
}

func TestJWTService(t *testing.T) {
	ts := new(JWTServiceTestSuite)
	tp, _ := filepath.Abs("../../../test/")
	ts.redisMock = new(mocks.MockRedisService)
	ts.rsaService = services.NewRSAService(tp+"/key.pem", tp+"/key.pem.pub")
	ts.rsaService.GenerateKeyPair()
	ts.rsaService.Write()
	ts.jwtService = services.NewJWTServiceRSA(
		ts.redisMock,
		ts.rsaService,
		services.JwtConfig{
			Exp:        15,
			ExpRefresh: 60,
			Keys: services.JwtKeys{
				Private: tp + "/key.pem",
				Public:  tp + "/key.pem.pub",
			},
		},
	)

	suite.Run(t, ts)

	os.Remove(tp + "/key.pem")
	os.Remove(tp + "/key.pem.pub")
}
