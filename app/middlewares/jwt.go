package middlewares

import (
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	Secret    string
	ExpSecret int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.Secret),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return baseResponse.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userID int) (string, string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtConf.ExpSecret) * time.Hour).Unix()

	unixTimeUTC := time.Unix(expirationTime, 0)

	unitTimeInRFC3339 := unixTimeUTC.UTC().Format(time.RFC3339)

	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Create token with claims

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.Secret))

	return token, unitTimeInRFC3339, err
}

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
