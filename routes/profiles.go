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
		Data: *convertedCities,
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
		Data: convertedProfile.Init(),
	}
	return c.Status(200).JSON(response)
}

