package tokenutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labbs/zotion/pkg/config"
	"github.com/labbs/zotion/pkg/models"
)

func CreateAccessToken(user_id, sessionId string) (accessToken string, err error) {
	exp := time.Now().Add(time.Second * time.Duration(config.Session.Expire)).Unix()
	claims := &models.JwtCustomClaims{
		SessionId: sessionId,
		UserId:    user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Session.Issuer,
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(exp, 0)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Session.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CreateRefreshToken(user_id, session_id string) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(config.Session.Expire)).Unix()
	claimsRefresh := &models.JwtCustomRefreshClaims{
		SessionId: session_id,
		UserId:    user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Session.Issuer,
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(exp, 0)},
		},
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	t, err := tokenRefresh.SignedString([]byte(config.Session.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetSessionInformationFromToken(token string) (user_id, sessionId string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Session.SecretKey), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		sessionId = claims["session_id"].(string)
		user_id = claims["user_id"].(string)
		return user_id, sessionId, nil
	}

	return "", "", fmt.Errorf("invalid token")
}

func IsAuthorized(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Session.SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// IsPasswordComplex checks if the password meets complexity requirements.
func IsPasswordComplex(password string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case (char >= '!' && char <= '/') || (char >= ':' && char <= '@') || (char >= '[' && char <= '`') || (char >= '{' && char <= '~'):
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

// IsSessionExpired checks if the session has expired based on the configured expiration time.
func IsSessionExpired(expire time.Time) bool {
	return expire.Before(time.Now())
}
