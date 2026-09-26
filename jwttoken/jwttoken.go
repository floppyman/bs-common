package jwttoken

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken creates a new JWT access token signed with HS512.
func GenerateAccessToken(secret string, issuer string, userID int, expiryMinutes int) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub": strconv.Itoa(userID),
		"iss": issuer,
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(expiryMinutes) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(secret))
}

// VerifyAccessToken parses and verifies a JWT access token, returning the user ID from the 'sub' claim.
func VerifyAccessToken(secret string, issuer string, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS512 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}, jwt.WithIssuer(issuer))
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims type")
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(sub)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
