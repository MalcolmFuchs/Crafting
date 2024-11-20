package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string, secretKey []byte) (string, error) {
	claims := AccessTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    "backend-crafting",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("ERROR: Fehler beim Signieren des Access Tokens für user_id %s: %v", userID, err)
		return "", err
	}

	log.Printf("INFO: Access Token erfolgreich generiert für user_id %s", userID)

	return signedToken, nil
}

func GenerateRefreshToken(userID string, secretKey []byte) (string, error) {
	claims := RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * 24 * time.Hour)),
			Issuer:    "backend-crafting",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("ERROR: Fehler beim Signieren des Acess Tokens für user_id %s: %v", userID, err)
		return "", err
	}

	log.Printf("INFO: Access Token erfolgreich generiert für user_id %s", userID)

	return signedToken, nil
}

func VerifyAccessToken(tokenString string, secretKey []byte) (*AccessTokenClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("ERROR: Unerwartetes Signaturverfahren: %v", t.Header["alg"])
			return nil, errors.New("unerwartetes Signaturverfahren")
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("ERROR: Fehler beim Parsen des Tokens: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		if claims.ExpiresAt.Time.Before(time.Now()) {
			log.Printf("ERROR: Access Token ist abgelaufen für user_id %s", claims.UserID)
			return nil, errors.New("access token ist abgelaufen")
		}
		log.Printf("INFO: Access Token erfolgreich verifiziert für user_id %s", claims.UserID)
		return claims, nil
	} else {
		log.Printf("ERROR: Ungültige Token-Claims")
		return nil, errors.New("ungültige token-claims")
	}
}
