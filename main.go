package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"laundry-hub-api/core"
	machineDeps "laundry-hub-api/machine/infrastructure/dependencies"
	machineRoutes "laundry-hub-api/machine/infrastructure/routes"
	userDeps "laundry-hub-api/user/infrastructure/dependencies"
	userRoutes "laundry-hub-api/user/infrastructure/routes"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Inicializar base de datos
	db := core.InitDB()
	defer db.Close()

	// Configurar Gin
	router := gin.Default()

	// Middleware CORS
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

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "API is running",
		})
	})

	// API Group
	api := router.Group("/api")

	// Inicializar dependencias de User
	userDependencies := userDeps.NewUserDependencies(db)

	// Inicializar dependencias de Machine
	machineDependencies := machineDeps.NewMachineDependencies(db)

	// Registrar rutas
	userRoutes.RegisterUserRoutes(api, userDependencies)
	machineRoutes.RegisterMachineRoutes(api, machineDependencies)

	// Obtener puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Iniciar servidor
	log.Printf("🚀 Server running on port %s", port)
	log.Printf("📡 API available at http://localhost:%s/api", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}