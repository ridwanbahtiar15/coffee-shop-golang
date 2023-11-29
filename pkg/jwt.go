package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type claims struct {
	Users_id   string    `json:"id"`
	Roles_id 	 string `json:"role"`
	jwt.RegisteredClaims
}

func NewPayload(users_id string, roles_id string, ) *claims {
	return &claims{
		Users_id:  users_id,
		Roles_id: roles_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
		},
	}
}

func (c *claims) GenerateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	result, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return result, err
}

func VerifyToken(token string) (*claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	payload := parsedToken.Claims.(*claims)
	return payload, nil
}