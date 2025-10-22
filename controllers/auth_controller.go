package controllers

import (
	"net/http"
	"time"

	"github.com/dawwasinha/ergovest-backend/config"
	"github.com/dawwasinha/ergovest-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Login request body struct
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate against user service
	su, err := services.ValidateCredentials(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  su.Username,
		"role": su.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	secret := []byte(config.GetJWTSecret())
	tokenString, err := token.SignedString(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful!",
		"token":   tokenString,
		"role":    su.Role,
	})
}
