package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	midw "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/sockets"
)

type Endpoint struct {
	DB *ent.Client
}

// All Endpoints (50)
func SetupRoutes(app *fiber.App, db *ent.Client) {
	endpoint := Endpoint{DB: db}
	midw := midw.Middleware{DB: db}

	api := app.Group("/api/v4")

	// HealthCheck Route (1)
	api.Get("/healthcheck", HealthCheck)

	// General Routes (1)
	generalRouter := api.Group("/general")
	generalRouter.Get("/site-detail", endpoint.GetSiteDetails)

	// Auth Routes (9)
	authRouter := api.Group("/auth")
	authRouter.Post("/register", endpoint.Register)
	authRouter.Post("/verify-email", endpoint.VerifyEmail)
	authRouter.Post("/resend-verification-email", endpoint.ResendVerificationEmail)
	authRouter.Post("/send-password-reset-otp", endpoint.SendPasswordResetOtp)
	authRouter.Post("/set-new-password", endpoint.SetNewPassword)
	authRouter.Post("/login", endpoint.Login)
	authRouter.Post("/refresh", endpoint.Refresh)
	authRouter.Get("/logout", midw.AuthMiddleware, endpoint.Logout)

	// Feed Routes (18)
	feedRouter := api.Group("/feed")
	feedRouter.Get("/posts", endpoint.RetrievePosts)
	feedRouter.Post("/posts", midw.AuthMiddleware, endpoint.CreatePost)
	feedRouter.Get("/posts/:slug", endpoint.RetrievePost)
	feedRouter.Put("/posts/:slug", midw.AuthMiddleware, endpoint.UpdatePost)
	feedRouter.Delete("/posts/:slug", midw.AuthMiddleware, endpoint.DeletePost)
	feedRouter.Get("/reactions/:focus/:slug", endpoint.RetrieveReactions)
	feedRouter.Post("/reactions/:focus/:slug", midw.AuthMiddleware, endpoint.CreateReaction)
	feedRouter.Delete("/reactions/:id", midw.AuthMiddleware, endpoint.DeleteReaction)
	feedRouter.Get("/posts/:slug/comments", endpoint.RetrieveComments)
	feedRouter.Post("/posts/:slug/comments", midw.AuthMiddleware, endpoint.CreateComment)
	feedRouter.Get("/comments/:slug", endpoint.RetrieveCommentWithReplies)
	feedRouter.Post("/comments/:slug", midw.AuthMiddleware, endpoint.CreateReply)
	feedRouter.Put("/comments/:slug", midw.AuthMiddleware, endpoint.UpdateComment)
	feedRouter.Delete("/comments/:slug", midw.AuthMiddleware, endpoint.DeleteComment)
	feedRouter.Get("/replies/:slug", endpoint.RetrieveReply)
	feedRouter.Put("/replies/:slug", midw.AuthMiddleware, endpoint.UpdateReply)
	feedRouter.Delete("/replies/:slug", midw.AuthMiddleware, endpoint.DeleteReply)

	// Profiles Routes (12)
	profilesRouter := api.Group("/profiles")
	profilesRouter.Get("/cities", endpoint.RetrieveCities)
	profilesRouter.Get("", midw.GuestMiddleware, endpoint.RetrieveUsers)
	profilesRouter.Get("/profile/:username", endpoint.RetrieveUserProfile)
	profilesRouter.Patch("/profile", midw.AuthMiddleware, endpoint.UpdateProfile)
	profilesRouter.Post("/profile", midw.AuthMiddleware, endpoint.DeleteUser)
	profilesRouter.Get("/friends", midw.AuthMiddleware, endpoint.RetrieveFriends)
	profilesRouter.Get("/friends/requests", midw.AuthMiddleware, endpoint.RetrieveFriendRequests)
	profilesRouter.Post("/friends/requests", midw.AuthMiddleware, endpoint.SendOrDeleteFriendRequest)
	profilesRouter.Put("/friends/requests", midw.AuthMiddleware, endpoint.AcceptOrRejectFriendRequest)
	profilesRouter.Get("/notifications", midw.AuthMiddleware, endpoint.RetrieveUserNotifications)
	profilesRouter.Post("/notifications", midw.AuthMiddleware, endpoint.ReadNotification)

	// Chat Routes (9)
	chatRouter := api.Group("/chats", midw.AuthMiddleware)
	chatRouter.Get("", endpoint.RetrieveUserChats)
	chatRouter.Post("", endpoint.SendMessage)
	chatRouter.Get("/:chat_id", endpoint.RetrieveMessages)
	chatRouter.Patch("/:chat_id", endpoint.UpdateGroupChat)
	chatRouter.Delete("/:chat_id", endpoint.DeleteGroupChat)
	chatRouter.Put("/messages/:message_id", endpoint.UpdateMessage)
	chatRouter.Delete("/messages/:message_id", endpoint.DeleteMessage)
	chatRouter.Post("/groups/group", endpoint.CreateGroupChat)
}

func SetupSockets(app *fiber.App) {
	app.Get("/api/v4/ws/notifications", websocket.New(sockets.NotificationSocket))
	app.Get("/api/v4/ws/chats/:id", websocket.New(sockets.ChatSocket))
}
