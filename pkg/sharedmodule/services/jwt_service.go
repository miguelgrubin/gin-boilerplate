package services

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTData struct {
	Jti    string `json:"jti"`
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

type JWTService interface {
	GenerateTokens(userId string, role string) (string, string)
	ValidateToken(token string) bool
	RefreshToken(token string, userId string, role string) (string, string, error)
	DecodeToken(token string) (JWTData, error)
	InvalidateToken(token string) error
}

type JWTServiceRSA struct {
	config       JwtConfig
	redisService RedisServiceInterface
}

var _ JWTService = &JWTServiceRSA{}

func NewJWTServiceRSA(rs RedisServiceInterface, c JwtConfig) *JWTServiceRSA {
	return &JWTServiceRSA{
		config:       c,
		redisService: rs,
	}
}

func (js *JWTServiceRSA) GenerateTokens(userId string, role string) (string, string) {
	refreshJti := uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"jti":     refreshJti,
		"user_id": userId,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * time.Duration(js.config.Exp)).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(js.config.Keys.Private))

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"jti":     refreshJti,
		"user_id": userId,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * time.Duration(js.config.ExpRefresh)).Unix(),
	})
	refreshTokenString, _ := refreshToken.SignedString([]byte(js.config.Keys.Private))

	js.redisService.Set(refreshJti, userId, time.Minute*time.Duration(js.config.ExpRefresh))
	return tokenString, refreshTokenString
}

func (js *JWTServiceRSA) ValidateToken(tokenString string) bool {
	_, err := js.DecodeToken(tokenString)
	if err != nil {
		return false
	}
	return true
}

func (js *JWTServiceRSA) RefreshToken(tokenString string, userId string, role string) (string, string, error) {
	data, err := js.DecodeToken(tokenString)

	if err != nil {
		return "", "", err
	}

	jti := data.Jti
	if jti == "" {
		return "", "", err
	}

	exists, err := js.redisService.Has(jti)
	if err != nil || !exists {
		return "", "", err
	}

	if data.UserId != userId {
		return "", "", errors.New("Invalid refresh token")
	}

	newTokenString, refreshTokenString := js.GenerateTokens(userId, role)
	js.redisService.Del(jti)

	return newTokenString, refreshTokenString, nil
}

func (js *JWTServiceRSA) DecodeToken(tokenString string) (JWTData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.config.Keys.Public), nil
	})
	if err != nil || !token.Valid {
		return JWTData{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return JWTData{}, err
	}
	jti, ok := claims["jti"].(string)
	if !ok {
		return JWTData{}, err
	}
	userId, ok := claims["user_id"].(string)
	if !ok {
		return JWTData{}, err
	}
	role, ok := claims["role"].(string)
	if !ok {
		return JWTData{}, err
	}
	return JWTData{
		Jti:    jti,
		UserId: userId,
		Role:   role,
	}, nil
}

func (js *JWTServiceRSA) InvalidateToken(tokenString string) error {
	data, err := js.DecodeToken(tokenString)
	if err != nil {
		return err
	}
	exists, err := js.redisService.Has(data.Jti)
	if err != nil || !exists {
		return err
	}
	err = js.redisService.Del(data.Jti)
	return err
}
