package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func getUser(c *fiber.Ctx, token string, db *ent.Client) (*ent.User, *string) {
	if len(token) < 8 {
		err := "Auth Token is Invalid or Expired!"
		return nil, &err
	}
	user, err := DecodeAccessToken(token[7:], db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	db := c.Locals("db").(*ent.Client)

	if len(token) < 1 {
		return c.Status(401).JSON(utils.ErrorResponse{Message: "Unauthorized User!"}.Init())
	}
	user, err := getUser(c, token, db)
	if err != nil {
		return c.Status(401).JSON(utils.ErrorResponse{Message: *err}.Init())
	}
	c.Locals("user", user)
	return c.Next()
}

func parseUUID(input string) *uuid.UUID {
    uuidVal, err := uuid.Parse(input)
	if err != nil {
		return nil
	}
    return &uuidVal
}