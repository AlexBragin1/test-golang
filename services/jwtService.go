package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"test/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenService struct {
	cryptoKey []byte
}

func NewJWTTokenService(secret []byte) *JWTTokenService {
	return &JWTTokenService{cryptoKey: secret}
}

func (s *JWTTokenService) GenerateRegToken(ctx context.Context, login domain.Login) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"exp":   jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		"flow":  domain.FLOW_REGISTRATION,
		"login": login,
	})

	tokenString, err := token.SignedString(s.cryptoKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTTokenService) GenerateAuthToken(ctx context.Context, user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"exp":     jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		"flow":    domain.FLOW_AUTHORIZATION,
		"user_id": user.ID,
		"groups":  fmt.Sprintf("%d", user.Groups),
	})

	tokenString, err := token.SignedString(s.cryptoKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTTokenService) getTokenString(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Token")
	if tokenString == "" {
		return "", errors.New("token header is empty or absent")
	}
	return tokenString, nil
}

func (s *JWTTokenService) parseTokenString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.cryptoKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, err
}

func (s *JWTTokenService) ReadFromToken(ctx context.Context, tokenString string, keys ...string) (map[string]string, error) {
	token, err := s.parseTokenString(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	res := make(map[string]string)

	for _, k := range keys {
		val, ok := claims[k]
		if !ok {
			res[k] = ""
			continue
		}
		strval, ok := val.(string)
		if ok {
			res[k] = strval
			continue
		}
		res[k] = ""
	}

	return res, nil
}

func (s *JWTTokenService) ReadFromRequest(r *http.Request, keys ...string) (map[string]string, error) {
	tokenString, err := s.getTokenString(r)
	if err != nil {
		return nil, err
	}
	strMap, err := s.ReadFromToken(r.Context(), tokenString, keys...)
	if err != nil {
		return nil, err
	}
	return strMap, nil
}
