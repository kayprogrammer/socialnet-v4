package main

import (
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/database"
	"github.com/kayprogrammer/socialnet-v4/initials"
	"github.com/kayprogrammer/socialnet-v4/routes"

	_ "github.com/kayprogrammer/socialnet-v4/docs"
)

// @title SOCIALNET API
// @version 4.0
// @description `A Realtime Social Networking API built with FIBER & GORM ORM` |
// @description
// @description `WEBSOCKETS:`
// @description  `Notifications:`
// @description `		URL: wss://{host}/api/v4/ws/notifications`
// @description `		* Requires authorization, so pass in the Bearer Authorization header.`
// @description `		* You can only read and not send notification messages into this socket.`
// @description `	Chats:`
// @description `		URL: wss://{host}/api/v4/ws/chats/{id}/`
// @description `		* Requires authorization, so pass in the Bearer Authorization header.`
// @description `		* Use chat_id as the ID for existing chat or username if its the first message in a DM.`
// @description `		* You cannot read realtime messages from a username that doesn't belong to the authorized user, but you can surely send messages.`
// @description `		* Only send message to the socket endpoint after the message has been created or updated, and files has been uploaded.`
// @description `		* Fields when sending message through the socket: e.g {"status": "CREATED", "id": "fe4e0235-80fc-4c94-b15e-3da63226f8ab"}`
// @description `			* status - This must be either CREATED or UPDATED (string type)`
// @description `			* id - This is the ID of the message (uuid type)`
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
	db := database.ConnectDb()
	initials.CreateInitialData(db)

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
		Title:    "SOCIALNET API Documentation",
	}

	app.Use(swagger.New(swaggerCfg))

	// Register Routes & Sockets
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	routes.SetupRoutes(app)
	routes.SetupSockets(app)
	log.Fatal(app.Listen(":8000"))
}
