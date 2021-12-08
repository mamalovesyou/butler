package jwt

import (
	"errors"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Access token are valid for 3 hour
const DefaultAccessTokenDuration = time.Hour * 3

// Refresh token are valid 30 days
const DefaultRefreshTokenDuration = time.Hour * 24 * 30

type AccessTokenClaims struct {
	Email     string
	FirstName string
	LastName  string
	TokenID   string
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	TokenID string
	jwt.StandardClaims
}

type JWTManager struct {
	secretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey:            secretKey,
		AccessTokenDuration:  DefaultAccessTokenDuration,
		RefreshTokenDuration: DefaultRefreshTokenDuration,
	}
}

func (m *JWTManager) NewAccessTokenClaims(sub, email, firtname, lastname string) *AccessTokenClaims {
	return &AccessTokenClaims{
		Email:     email,
		FirstName: firtname,
		LastName:  lastname,
		TokenID:   uuid.New().String(),
		StandardClaims: jwt.StandardClaims{
			Subject: sub,
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(m.AccessTokenDuration).Unix(),
		},
	}
}

func (m *JWTManager) CreateAccessToken(claims *AccessTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) NewRefreshTokenClaims(sub string) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		TokenID: uuid.New().String(),
		StandardClaims: jwt.StandardClaims{
			Subject: sub,
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(m.RefreshTokenDuration).Unix(),
		},
	}
}

func (m *JWTManager) CreateRefreshToken(claims *RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) ParseToken(token string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return jwt.MapClaims{}, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})
}

func (m *JWTManager) IsValidAccessToken(token string) (*AccessTokenClaims, error) {
	accessToken, err := m.ParseToken(token, &AccessTokenClaims{})
	if err != nil {
		return &AccessTokenClaims{}, err
	}

	if claims, ok := accessToken.Claims.(*AccessTokenClaims); ok && accessToken.Valid {
		return claims, nil
	}

	return &AccessTokenClaims{}, errors.New("Unable to parse claims")
}

func (m *JWTManager) IsValidRefreshToken(token string) (*RefreshTokenClaims, error) {
	refreshToken, err := m.ParseToken(token, &RefreshTokenClaims{})
	if err != nil {
		return &RefreshTokenClaims{}, err
	}

	if claims, ok := refreshToken.Claims.(*RefreshTokenClaims); ok && refreshToken.Valid {
		return claims, nil
	}

	return &RefreshTokenClaims{}, errors.New("Unable to parse claims")
}
