package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type DefaultMiddleware interface {
	JwtAuth() gin.HandlerFunc
}

type MyClaims struct {
	jwt.StandardClaims
	UserId   string `json:"UserId"`
	Username string `json:"Username"`
}

type defaultMiddleware struct{}

func NewMiddleware() DefaultMiddleware {
	return &defaultMiddleware{}
}

func (m *defaultMiddleware) JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		authorizationSplited := strings.Split(authorizationHeader, " ")
		if len(authorizationSplited) != 2 {
			fmt.Println("auth != 2")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if authorizationSplited[0] != "Bearer" {
			fmt.Println("auth != Bearer")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(authorizationSplited[1], func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
		})

		if err != nil {
			fmt.Println("parse ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("claims ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}


		c.Set("userInfo", claims)

		c.Next()

	}

}
