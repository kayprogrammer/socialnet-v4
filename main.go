package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/database"
	"github.com/kayprogrammer/socialnet-v4/initials"
	"github.com/kayprogrammer/socialnet-v4/routes"

	_ "github.com/kayprogrammer/socialnet-v4/docs"
)

// @title SOCIALNET API
// @version 4.0
// @description A simple bidding API built with Fiber
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
	db := database.Database.Db
	initials.CreateInitialData(db)

	app := fiber.New()

	// Set up the database middleware
	app.Use(database.DatabaseMiddleware)

	// CORS config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Guestuserid, Access-Control-Allow-Origin, Content-Disposition",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	// Inject environment text
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("env", "normal")
		return c.Next()
	})

	// Register routes
	routes.SetupRoutes(app)
	app.Get("/*", swagger.HandlerDefault) // default

	log.Fatal(app.Listen(":8000"))
}
