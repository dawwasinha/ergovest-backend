package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dawwasinha/ergovest-backend/services" // Ganti dengan path modul Anda
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func GetSurveyStats(c *gin.Context) {
	// TODO: Implementasi logika kalkulasi statistik survei
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint GetSurveyStats not yet implemented"})
}

func SubmitSurvey(c *gin.Context) {
	// TODO: Implementasi logika validasi dan penyimpanan survei
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint SubmitSurvey not yet implemented"})
}

func WebSocketHandler(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("ws upgrade error: %v", err)
		return
	}

	client := services.RegisterWSClient(conn)

	// Keep reading to detect close
	for {
		if _, _, err := conn.NextReader(); err != nil {
			services.UnregisterWSClient(client)
			break
		}
	}
}
