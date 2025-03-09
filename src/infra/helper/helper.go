package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// TokenClaims menyimpan klaim JWT
type TokenClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken membuat token JWT
func GenerateToken(userID int64, email string, isRefreshToken bool) (string, error) {
	var expirationTime time.Time

	// Set expiration time berdasarkan jenis token
	if isRefreshToken {
		// Refresh token berlaku selama 14 hari
		expirationTime = time.Now().Add(14 * 24 * time.Hour)
	} else {
		// Access token berlaku selama  1 hari
		expirationTime = time.Now().Add(1 * 24 * time.Hour)
	}

	claims := &TokenClaims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Buat JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	return token.SignedString(jwtKey)
}

// VerifyToken memverifikasi token JWT
func VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

const (
	Page    = int64(1)
	PerPage = int64(10)
)
