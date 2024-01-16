package routes

import (
	"github.com/gofiber/fiber/v2"
	midw "github.com/kayprogrammer/socialnet-v4/authentication"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v4")

	// HealthCheck Route
	api.Get("/healthcheck", HealthCheck)

	// General Routes
	generalRouter := api.Group("/general")
	generalRouter.Get("/site-detail", GetSiteDetails)

	// Auth Routes
	authRouter := api.Group("/auth")
	authRouter.Post("/register", Register)
	authRouter.Post("/verify-email", VerifyEmail)
	authRouter.Post("/resend-verification-email", ResendVerificationEmail)
	authRouter.Post("/send-password-reset-otp", SendPasswordResetOtp)
	authRouter.Post("/set-new-password", SetNewPassword)
	authRouter.Post("/login", Login)
	authRouter.Post("/refresh", Refresh)
	authRouter.Get("/logout", midw.AuthMiddleware, Logout)

	// Feed Routes
	feedRouter := api.Group("/feed")
	feedRouter.Get("/posts", RetrievePosts)
	feedRouter.Post("/posts", midw.AuthMiddleware, CreatePost)
	feedRouter.Get("/posts/:slug", RetrievePost)
	feedRouter.Put("/posts/:slug", midw.AuthMiddleware, UpdatePost)
	feedRouter.Delete("/posts/:slug", midw.AuthMiddleware, DeletePost)
}
