package main

import (
	"log"

	"github.com/dawwasinha/ergovest-backend/config"
	"github.com/dawwasinha/ergovest-backend/routes"
	"github.com/dawwasinha/ergovest-backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inisialisasi Database (Contoh)
	// db.ConnectDB()
	services.InitDB()
	defer services.CloseDB()
	services.InitUsers()

	// 2. Inisialisasi MQTT Service (Koneksi ke HiveMQ di background)
	// 'go' menjalankan StartMQTTClient() secara asynchronous
	go services.StartMQTTClient() // <-- 2. Panggil service MQTT dalam goroutine

	// 3. Setup Gin Router
	router := gin.Default()

	// CORS Middleware (Penting agar React bisa mengakses)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 4. Daftarkan Semua Routes (API Endpoints)
	routes.SetupRouter(router)

	// Serve API docs (ReDoc) from /docs
	router.Static("/docs", "./docs")

	// 5. Jalankan Server Gin
	log.Fatal(router.Run(config.GetServerPort())) // Server berjalan pada port dari env
}
