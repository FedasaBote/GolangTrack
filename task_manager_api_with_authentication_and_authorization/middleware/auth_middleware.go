package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


var jwtSecret = []byte(getJwtSecret("JWT_SECRET"))

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface {}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID, userIDExists := claims["id"].(string)
            userRole, userRoleExists := claims["role"].(string)
			if !userIDExists || !userRoleExists {
                c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT claims"})
                c.Abort()
                return
            }

            c.Set("userID", userID)
            c.Set("userRole", userRole)
        } else {
            c.IndentedJSON(http.StatusUnauthorized, "Invalid JWT claims")
            c.Abort()
            return
        }

		c.Next()
	}
}

func RoleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("userRole")
        if !exists || userRole != "admin" {
            c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}


func getJwtSecret(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}

	return os.Getenv(key)
}