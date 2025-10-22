package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/dawwasinha/ergovest-backend/routes"
	"ergovest-backend/services" // <-- 1. Import package services
)

func main() {
	// 1. Inisialisasi Database (Contoh)
	// db.ConnectDB()

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

	// 5. Jalankan Server Gin
	log.Fatal(router.Run(":8080")) // Server berjalan di port 8080
}