package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/dawwasinha/ergovest-backend/config"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
            c.Abort()
            return
        }

        // Expect Bearer <token>
        var tokenString string
        if len(auth) > 7 && auth[:7] == "Bearer " {
            tokenString = auth[7:]
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
            c.Abort()
            return
        }

        secret := []byte(config.GetJWTSecret())
        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return secret, nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
