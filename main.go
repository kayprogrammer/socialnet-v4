package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/contrib/swagger"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/database"
	"github.com/kayprogrammer/socialnet-v4/routes"

	_ "github.com/kayprogrammer/socialnet-v4/docs"
)

// @title SOCIALNET API
// @version 4.0
// @description A Realtime Social Networking API built with Fiber
// @Accept json
// @Produce json
// @BasePath  /api/v4
// @Security BearerAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type 'Bearer jwt_string' to correctly set the API Key
func main() {
	cfg := config.GetConfig()
	database.ConnectDb()
	// initials.CreateInitialData(db)

	app := fiber.New()

	// Set up the database middleware
	app.Use(database.DatabaseMiddleware)

	// CORS config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Access-Control-Allow-Origin, Content-Disposition",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	// Inject environment text
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("env", "normal")
		return c.Next()
	})

	// Swagger Config
	swaggerCfg := swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "/",
		Title: "SOCIALNET API Documentation",
	}
	
	app.Use(swagger.New(swaggerCfg))

	// Register Routes & Sockets
	app.Use("/ws", func(c *fiber.Ctx) error {
        // IsWebSocketUpgrade returns true if the client
        // requested upgrade to the WebSocket protocol.
        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    })

	routes.SetupRoutes(app)
	routes.SetupSockets(app)
	log.Fatal(app.Listen(":8000"))
}
