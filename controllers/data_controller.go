package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/dawwasinha/ergovest-backend/services" // Ganti dengan path modul Anda
)

// GetSensorHistory mengambil data sensor historis untuk grafik
func GetSensorHistory(c *gin.Context) {
	// Ambil parameter limit (berapa banyak data yang diminta React)
	limitStr := c.DefaultQuery("limit", "100") 
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100 // Default jika parameter salah
	}

	history := services.GetSensorHistory(limit)
	c.JSON(http.StatusOK, history)
}

// GetAlertHistory mengambil data riwayat alert
func GetAlertHistory(c *gin.Context) {
	history := services.GetAlertHistory()
	c.JSON(http.StatusOK, history)
}

// TODO: Tambahkan fungsi untuk SubmitSurvey dan GetSurveyStats di sini