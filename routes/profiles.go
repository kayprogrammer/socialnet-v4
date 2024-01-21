package routes

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var cityManager = managers.CityManager{}

// @Summary Retrieve cities based on query params
// @Description This endpoint retrieves the first 10 cities that matches the query params
// @Tags Profiles
// @Param name query string false "City name"
// @Success 200 {object} schemas.CitiesResponseSchema
// @Router /profiles/cities [get]
func RetrieveCities(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	message := "Cities Fetched"
	name := c.Query("name")

	// Define a regular expression to match non-word characters (excluding spaces).
	re := regexp.MustCompile(`[^\w\s]`)
	// Use the regular expression to replace matching substrings with an empty string.
	name = re.ReplaceAllString(name, "")

	cities := cityManager.All(db, name)

	// Convert type and return Cities
	convertedCities := utils.ConvertStructData(cities, []schemas.CitySchema{}).(*[]schemas.CitySchema)
	if len(*convertedCities) == 0 {
		message = "No match found"
	}
	response := schemas.CitiesResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: message}.Init(),
		Data:           *convertedCities,
	}.Init()
	return c.Status(200).JSON(response)
}

var userProfileManager = managers.UserProfileManager{}

// @Summary Retrieve Users
// @Description This endpoint retrieves a paginated list of users
// @Tags Profiles
// @Param page query int false "Current Page" default(1)
// @Success 200 {object} schemas.ProfilesResponseSchema
// @Router /profiles [get]
// @Security BearerAuth
func RetrieveUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)

	users := userProfileManager.GetUsers(db, user)

	// Paginate, Convert type and return Users
	paginatedData, paginatedUsers, err := PaginateQueryset(users, c)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	convertedProfiles := utils.ConvertStructData(paginatedUsers, []schemas.ProfileSchema{}).(*[]schemas.ProfileSchema)
	response := schemas.ProfilesResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Users fetched"}.Init(),
		Data: schemas.ProfilesResponseDataSchema{
			PaginatedResponseDataSchema: *paginatedData,
			Items:                       *convertedProfiles,
		}.Init(),
	}
	return c.Status(200).JSON(response)
}

// @Summary Retrieve User Profile
// @Description This endpoint retrieves a user profile
// @Tags Profiles
// @Param username path string true "Username of user"
// @Success 200 {object} schemas.ProfileResponseSchema
// @Router /profiles/profile/{username} [get]
func RetrieveUserProfile(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	username := c.Params("username")

	user, errData := userProfileManager.GetByUsername(db, username)
	if errData != nil {
		return c.Status(404).JSON(errData)
	}

	// Convert type and return User
	convertedProfile := utils.ConvertStructData(user, schemas.ProfileSchema{}).(*schemas.ProfileSchema)
	response := schemas.ProfileResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "User details fetched"}.Init(),
		Data:           convertedProfile.Init(),
	}
	return c.Status(200).JSON(response)
}

// @Summary Update User Profile
// @Description This endpoint updates a user profile
// @Tags Profiles
// @Param profile body schemas.ProfileUpdateSchema true "Profile object"
// @Success 200 {object} schemas.ProfileResponseSchema
// @Router /profiles/profile [patch]
// @Security BearerAuth
func UpdateProfile(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)

	profileData := schemas.ProfileUpdateSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &profileData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(profileData); err != nil {
		return c.Status(422).JSON(err)
	}

	// Validate City Value
	cityID := profileData.CityID
	if cityID != nil {
		city := cityManager.GetByID(db, *cityID)
		if city == nil {
			data := map[string]string{
				"city_id": "No city with that ID",
			}
			return c.Status(422).JSON(utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid Entry", data))
		}
		profileData.City = city
	}

	updatedProfile := userProfileManager.Update(db, user, profileData)

	// Convert type and return User
	convertedProfile := utils.ConvertStructData(updatedProfile, schemas.ProfileUpdateResponseDataSchema{}).(*schemas.ProfileUpdateResponseDataSchema)
	response := schemas.ProfileUpdateResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "User updated fetched"}.Init(),
		Data:           convertedProfile.Init(profileData.FileType),
	}
	return c.Status(200).JSON(response)
}

// @Summary Delete User's Account
// @Description This endpoint deletes a particular user's account (irreversible)
// @Tags Profiles
// @Param password body schemas.DeleteUserSchema true "Password"
// @Success 200 {object} schemas.ResponseSchema
// @Router /profiles/profile [post]
// @Security BearerAuth
func DeleteUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)

	deleteUserData := schemas.DeleteUserSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &deleteUserData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(deleteUserData); err != nil {
		return c.Status(422).JSON(err)
	}

	// Check if password is valid
	if !utils.CheckPasswordHash(deleteUserData.Password, user.Password) {
		data := map[string]string{
			"password": "Incorrect password",
		}
		return c.Status(422).JSON(utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid Entry", data))
	}

	// Delete User
	db.User.DeleteOne(user).Exec(managers.Ctx)
	response := schemas.ResponseSchema{Message: "User deleted"}.Init()
	return c.Status(200).JSON(response)
}
