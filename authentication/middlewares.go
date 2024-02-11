package authentication

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type Middleware struct {
	DB *ent.Client
}

func GetUser(token string, db *ent.Client) (*ent.User, *string) {
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

func (mid Middleware) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	db := mid.DB

	if len(token) < 1 {
		return c.Status(401).JSON(utils.RequestErr(utils.ERR_UNAUTHORIZED_USER, "Unauthorized User!"))
	}
	user, err := GetUser(token, db)
	if err != nil {
		return c.Status(401).JSON(utils.RequestErr(utils.ERR_INVALID_TOKEN, *err))
	}
	c.Locals("user", user)
	return c.Next()
}

func (mid Middleware) GuestMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	db := mid.DB
	var user *ent.User
	if len(token) > 0 {
		userObj, err := GetUser(token, db)
		if err != nil {
			return c.Status(401).JSON(utils.RequestErr(utils.ERR_INVALID_TOKEN, *err))
		}
		user = userObj
	}
	c.Locals("user", user)
	return c.Next()
}
