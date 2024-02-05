package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	midw "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/sockets"
)

// All Endpoints (50)
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v4")

	// HealthCheck Route (1)
	api.Get("/healthcheck", HealthCheck)

	// General Routes (1)
	generalRouter := api.Group("/general")
	generalRouter.Get("/site-detail", GetSiteDetails)

	// Auth Routes (9)
	authRouter := api.Group("/auth")
	authRouter.Post("/register", Register)
	authRouter.Post("/verify-email", VerifyEmail)
	authRouter.Post("/resend-verification-email", ResendVerificationEmail)
	authRouter.Post("/send-password-reset-otp", SendPasswordResetOtp)
	authRouter.Post("/set-new-password", SetNewPassword)
	authRouter.Post("/login", Login)
	authRouter.Post("/refresh", Refresh)
	authRouter.Get("/logout", midw.AuthMiddleware, Logout)

	// Feed Routes (18)
	feedRouter := api.Group("/feed")
	feedRouter.Get("/posts", RetrievePosts)
	feedRouter.Post("/posts", midw.AuthMiddleware, CreatePost)
	feedRouter.Get("/posts/:slug", RetrievePost)
	feedRouter.Put("/posts/:slug", midw.AuthMiddleware, UpdatePost)
	feedRouter.Delete("/posts/:slug", midw.AuthMiddleware, DeletePost)
	feedRouter.Get("/reactions/:focus/:slug", RetrieveReactions)
	feedRouter.Post("/reactions/:focus/:slug", midw.AuthMiddleware, CreateReaction)
	feedRouter.Delete("/reactions/:id", midw.AuthMiddleware, DeleteReaction)
	feedRouter.Get("/posts/:slug/comments", RetrieveComments)
	feedRouter.Post("/posts/:slug/comments", midw.AuthMiddleware, CreateComment)
	feedRouter.Get("/comments/:slug", RetrieveCommentWithReplies)
	feedRouter.Post("/comments/:slug", midw.AuthMiddleware, CreateReply)
	feedRouter.Put("/comments/:slug", midw.AuthMiddleware, UpdateComment)
	feedRouter.Delete("/comments/:slug", midw.AuthMiddleware, DeleteComment)
	feedRouter.Get("/replies/:slug", RetrieveReply)
	feedRouter.Put("/replies/:slug", midw.AuthMiddleware, UpdateReply)
	feedRouter.Delete("/replies/:slug", midw.AuthMiddleware, DeleteReply)

	// Profiles Routes (12)
	profilesRouter := api.Group("/profiles")
	profilesRouter.Get("/cities", RetrieveCities)
	profilesRouter.Get("", midw.GuestMiddleware, RetrieveUsers)
	profilesRouter.Get("/profile/:username", RetrieveUserProfile)
	profilesRouter.Patch("/profile", midw.AuthMiddleware, UpdateProfile)
	profilesRouter.Post("/profile", midw.AuthMiddleware, DeleteUser)
	profilesRouter.Get("/friends", midw.AuthMiddleware, RetrieveFriends)
	profilesRouter.Get("/friends/requests", midw.AuthMiddleware, RetrieveFriendRequests)
	profilesRouter.Post("/friends/requests", midw.AuthMiddleware, SendOrDeleteFriendRequest)
	profilesRouter.Put("/friends/requests", midw.AuthMiddleware, AcceptOrRejectFriendRequest)
	profilesRouter.Get("/notifications", midw.AuthMiddleware, RetrieveUserNotifications)
	profilesRouter.Post("/notifications", midw.AuthMiddleware, ReadNotification)

	// Chat Routes (9)
	chatRouter := api.Group("/chats", midw.AuthMiddleware)
	chatRouter.Get("", RetrieveUserChats)
	chatRouter.Post("", SendMessage)
	chatRouter.Get("/:chat_id", RetrieveMessages)
	chatRouter.Patch("/:chat_id", UpdateGroupChat)
	chatRouter.Delete("/:chat_id", DeleteGroupChat)
	chatRouter.Put("/messages/:message_id", UpdateMessage)
	chatRouter.Delete("/messages/:message_id", DeleteMessage)
	chatRouter.Post("/groups/group", CreateGroupChat)
}

func SetupSockets(app *fiber.App) {
	app.Get("/api/v4/ws/notifications", websocket.New(sockets.NotificationSocket))
	app.Get("/api/v4/ws/chats/:id", websocket.New(sockets.ChatSocket))
}
