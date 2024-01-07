package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

// @Summary Retrieve site details
// @Description This endpoint retrieves few details of the site/application.
// @Tags General
// @Success 200 {object} schemas.SiteDetailResponseSchema
// @Router /general/site-detail [get]
func GetSiteDetails(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)

	sitedetail, _ := managers.SiteDetailManager{}.Get(db)
	if sitedetail == nil {
		sitedetail, _ = managers.SiteDetailManager{}.Create(db)
	}
	responseSiteDetail := schemas.SiteDetailResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Site Details Fetched!"}.Init(),
		Data:           sitedetail,
	}
	return c.Status(200).JSON(responseSiteDetail)
}
