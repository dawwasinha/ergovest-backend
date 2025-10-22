package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "ergovest-backend/models" // Jika Anda menggunakan model user
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

    // Replikasi logika validasi Node-RED di sini
    // var validUsers = {'admin': 'admin123', 'operator': 'operator123', 'supervisor': 'super123'};
    if (req.Username == "admin" && req.Password == "admin123") {
        // Pada aplikasi nyata: Generate JWT Token
        c.JSON(http.StatusOK, gin.H{
            "message": "Login Successful!",
            "token": "fake-jwt-token-admin", 
            "role": "admin",
        })
        return
    }

    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}