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