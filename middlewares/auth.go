package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/payslip/config"
)

var AppConfig *config.Config

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		AppConfig, err = config.LoadConfig()
		if err != nil {
			panic("Failed to load config: " + err.Error())
		}
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(AppConfig.AuthKey), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims["role"] != role {
			fmt.Println("role", claims["role"], role)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized role"})
			return
		}

		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("role", claims["role"].(string))
		c.Next()
	}
}

func GenerateToken(userID uint, role string) (string, error) {
	var err error
	AppConfig, err = config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(AppConfig.AuthKey))
}
