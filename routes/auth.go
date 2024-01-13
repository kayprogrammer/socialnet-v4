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

var userManager = managers.UserManager{}
var otpManager = managers.OtpManager{}

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
	otp := otpManager.GetOrCreate(db, newUser.ID)
	go senders.SendEmail(c.Locals("env"), newUser, "activate", &otp.Code)

	response := schemas.RegisterResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Registration successful"}.Init(),
		Data:           schemas.EmailRequestSchema{Email: newUser.Email},
	}
	return c.Status(201).JSON(response)
}

// @Summary Verify a user's email
// @Description This endpoint verifies a user's email.
// @Tags Auth
// @Param verify_email body schemas.VerifyEmailRequestSchema true "Verify Email object"
// @Success 200 {object} schemas.ResponseSchema
// @Failure 422 {object} utils.ErrorResponse
// @Router /auth/verify-email [post]
func VerifyEmail(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	validator := utils.Validator()

	verifyEmail := schemas.VerifyEmailRequestSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &verifyEmail); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(verifyEmail); err != nil {
		return c.Status(422).JSON(err)
	}

	user, _ := userManager.GetByEmail(db, verifyEmail.Email)
	if user == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_EMAIL, Message: "Incorrect Email"}.Init())
	}

	if user.IsEmailVerified {
		return c.Status(200).JSON(schemas.ResponseSchema{Message: "Email already verified"}.Init())
	}

	otp, _ := otpManager.GetByUserID(db, user.ID)
	if otp == nil || otp.Code != verifyEmail.Otp {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_OTP, Message: "Incorrect Otp"}.Init())
	}

	if otpManager.CheckExpiration(otp) {
		return c.Status(400).JSON(utils.ErrorResponse{Code: utils.ERR_EXPIRED_OTP, Message: "Expired Otp"}.Init())
	}

	// Update User
	user.Update().SetIsEmailVerified(true).Save(managers.Ctx)

	// Send Welcome Email
	go senders.SendEmail(c.Locals("env"), user, "welcome", nil)

	response := schemas.ResponseSchema{Message: "Account verification successful"}.Init()
	return c.Status(200).JSON(response)
}

// @Summary Resend Verification Email
// @Description This endpoint resends new otp to the user's email.
// @Tags Auth
// @Param email body schemas.EmailRequestSchema true "Email object"
// @Success 200 {object} schemas.ResponseSchema
// @Failure 422 {object} utils.ErrorResponse
// @Router /auth/resend-verification-email [post]
func ResendVerificationEmail(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	validator := utils.Validator()

	emailSchema := schemas.EmailRequestSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &emailSchema); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(emailSchema); err != nil {
		return c.Status(422).JSON(err)
	}

	user, _ := userManager.GetByEmail(db, emailSchema.Email)
	if user == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_EMAIL, Message: "Incorrect Email"}.Init())
	}

	if user.IsEmailVerified {
		return c.Status(200).JSON(schemas.ResponseSchema{Message: "Email already verified"}.Init())
	}

	// Send Email
	otp := otpManager.GetOrCreate(db, user.ID)
	go senders.SendEmail(c.Locals("env"), user, "activate", &otp.Code)

	response := schemas.ResponseSchema{Message: "Verification email sent"}.Init()
	return c.Status(200).JSON(response)
}

// @Summary Send Password Reset Otp
// @Description This endpoint sends new password reset otp to the user's email.
// @Tags Auth
// @Param email body schemas.EmailRequestSchema true "Email object"
// @Success 200 {object} schemas.ResponseSchema
// @Failure 422 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /auth/send-password-reset-otp [post]
func SendPasswordResetOtp(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	validator := utils.Validator()

	emailSchema := schemas.EmailRequestSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &emailSchema); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(emailSchema); err != nil {
		return c.Status(422).JSON(err)
	}

	user, _ := userManager.GetByEmail(db, emailSchema.Email)
	if user == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_EMAIL, Message: "Incorrect Email"}.Init())
	}

	// Send Email
	otp := otpManager.GetOrCreate(db, user.ID)
	go senders.SendEmail(c.Locals("env"), user, "reset", &otp.Code)

	response := schemas.ResponseSchema{Message: "Password otp sent"}.Init()
	return c.Status(200).JSON(response)
}

// @Summary Set New Password
// @Description This endpoint verifies the password reset otp.
// @Tags Auth
// @Param email body schemas.SetNewPasswordSchema true "Password reset object"
// @Success 200 {object} schemas.ResponseSchema
// @Failure 422 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /auth/set-new-password [post]
func SetNewPassword(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	validator := utils.Validator()

	passwordResetSchema := schemas.SetNewPasswordSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &passwordResetSchema); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(passwordResetSchema); err != nil {
		return c.Status(422).JSON(err)
	}

	user, _ := userManager.GetByEmail(db, passwordResetSchema.Email)
	if user == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_EMAIL, Message: "Incorrect Email"}.Init())
	}

	otp, _ := otpManager.GetByUserID(db, user.ID)
	if otp == nil || otp.Code != passwordResetSchema.Otp {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_INCORRECT_OTP, Message: "Incorrect Otp"}.Init())
	}

	if otpManager.CheckExpiration(otp) {
		return c.Status(400).JSON(utils.ErrorResponse{Code: utils.ERR_EXPIRED_OTP, Message: "Expired Otp"}.Init())
	}

	// Set Password
	user.Update().SetPassword(utils.HashPassword(passwordResetSchema.Password)).Save(managers.Ctx)

	// Send Email
	go senders.SendEmail(c.Locals("env"), user, "reset-success", nil)

	response := schemas.ResponseSchema{Message: "Password reset successful"}.Init()
	return c.Status(200).JSON(response)
}
