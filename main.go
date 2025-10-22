package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "ergovest-backend/routes" // Ganti 'ergovest-backend' dengan nama modul Anda

    // Tambahkan import service lain di sini
)

func main() {
    // 1. Inisialisasi Database (Contoh)
    // db.ConnectDB() 

    // 2. Inisialisasi MQTT Service (Koneksi ke HiveMQ di background)
    // go services.StartMQTTClient() 

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