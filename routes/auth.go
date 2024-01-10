package routes

import (
	"github.com/gofiber/fiber/v2"

	// auth "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/senders"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

// @Summary Register a new user
// @Description This endpoint registers new users into our application.
// @Tags Auth
// @Param user body schemas.RegisterUser true "User object"
// @Success 201 {object} schemas.RegisterResponseSchema
// @Failure 422 {object} utils.ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	validator := utils.Validator()

	user := schemas.RegisterUser{}
	userManager := managers.UserManager{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &user); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(user); err != nil {
		return c.Status(422).JSON(err)
	}

	userByEmail, _ := userManager.GetByEmail(db, user.Email)
	if userByEmail != nil {
		data := map[string]string{
			"email": "Email already registered!",
		}
		return c.Status(422).JSON(utils.ErrorResponse{Code: utils.ERR_INVALID_ENTRY, Message: "Invalid Entry", Data: &data}.Init())
	}

	// Create User
	newUser, _ := userManager.Create(db, user)

	// Send Email
	otp := managers.OtpManager{}.GetOrCreate(db, newUser.ID)
	go senders.SendEmail(c.Locals("env"), newUser, "activate", &otp.Code)

	response := schemas.RegisterResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Registration successful"}.Init(),
		Data:           schemas.EmailRequestSchema{Email: newUser.Email},
	}
	return c.Status(201).JSON(response)
}
