package services

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTData struct {
	Jti    string `json:"jti"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type JWTService interface {
	GenerateTokens(userID string, role string) (string, string, error)
	ValidateToken(token string) bool
	RefreshToken(token string, userID string, role string) (string, string, error)
	DecodeToken(token string) (JWTData, error)
	InvalidateToken(token string) error
}

type JWTServiceRSA struct {
	config       JwtConfig
	redisService RedisServiceInterface
	rsaService   RSAService
}

var _ JWTService = &JWTServiceRSA{}

func NewJWTServiceRSA(redis RedisServiceInterface, rsa RSAService, c JwtConfig) *JWTServiceRSA {
	return &JWTServiceRSA{
		config:       c,
		redisService: redis,
		rsaService:   rsa,
	}
}

func (js *JWTServiceRSA) GenerateTokens(userID string, role string) (string, string, error) {
	refreshJti := uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"jti":     refreshJti,
		"user_id": userID,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * time.Duration(js.config.Exp)).Unix(),
	})
	tokenString, err := token.SignedString(js.rsaService.GetPrivateKey())
	if err != nil {
		log.Println("error generating token for user:", userID, err)
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"jti":     refreshJti,
		"user_id": userID,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * time.Duration(js.config.ExpRefresh)).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString(js.rsaService.GetPrivateKey())
	if err != nil {
		log.Println("error generating refresh token for user:", userID, err)
		return "", "", err
	}

	js.redisService.Set(refreshJti, userID, time.Minute*time.Duration(js.config.ExpRefresh))
	return tokenString, refreshTokenString, nil
}

func (js *JWTServiceRSA) ValidateToken(tokenString string) bool {
	_, err := js.DecodeToken(tokenString)
	if err != nil {
		return false
	}
	return true
}

func (js *JWTServiceRSA) RefreshToken(tokenString string, userID string, role string) (string, string, error) {
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

	if data.UserID != userID {
		return "", "", errors.New("Invalid refresh token")
	}

	newTokenString, refreshTokenString, err := js.GenerateTokens(userID, role)
	if err != nil {
		return "", "", err
	}
	js.redisService.Del(jti)

	return newTokenString, refreshTokenString, nil
}

func (js *JWTServiceRSA) DecodeToken(tokenString string) (JWTData, error) {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (any, error) {
		return js.rsaService.GetPublicKey(), nil
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
	userID, ok := claims["user_id"].(string)
	if !ok {
		return JWTData{}, err
	}
	role, ok := claims["role"].(string)
	if !ok {
		return JWTData{}, err
	}
	return JWTData{
		Jti:    jti,
		UserID: userID,
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
