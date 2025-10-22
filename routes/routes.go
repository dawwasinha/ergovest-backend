package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/dawwasinha/ergovest-backend/controllers" // Ganti dengan path modul Anda
)
    
func SetupRouter(r *gin.Engine) {
    // Group untuk API v1 (API yang diakses oleh React)
    v1 := r.Group("/api/v1")
    {
        // 1. Authentication (Login)
        v1.POST("/login", controllers.Login)

        // 2. Data Endpoint (Membutuhkan otorisasi setelah login)
        dataRoutes := v1.Group("/data")
        // dataRoutes.Use(middlewares.AuthMiddleware()) // Middleware JWT untuk proteksi
        {
            // Untuk data historis dan statistik
            dataRoutes.GET("/sensor-history", controllers.GetSensorHistory)
            dataRoutes.GET("/alert-history", controllers.GetAlertHistory)
            dataRoutes.GET("/survey-stats", controllers.GetSurveyStats)
        }
        
        // 3. Survey Endpoint
        v1.POST("/survey/submit", controllers.SubmitSurvey)

        // 4. WebSocket Endpoint (Untuk data real-time)
        v1.GET("/ws", controllers.WebSocketHandler)
    }

    // Hello World Gin Test
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ErgoVest Backend API is running!",
        })
    })
}