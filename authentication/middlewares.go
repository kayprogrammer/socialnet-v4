package authentication

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func getUser(c *fiber.Ctx, token string, db *ent.Client) (*ent.User, *string) {
	if !strings.HasPrefix(token, "Bearer ") {
		err := "Auth Bearer Not Provided"
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
		return c.Status(401).JSON(utils.ErrorResponse{Code: utils.ERR_UNAUTHORIZED_USER, Message: "Unauthorized User!"}.Init())
	}
	user, err := getUser(c, token, db)
	if err != nil {
		return c.Status(401).JSON(utils.ErrorResponse{Code: utils.ERR_INVALID_TOKEN, Message: *err}.Init())
	}
	c.Locals("user", user)
	return c.Next()
}

func ParseUUID(input string) *uuid.UUID {
	uuidVal, err := uuid.Parse(input)
	if err != nil {
		return nil
	}
	return &uuidVal
}
