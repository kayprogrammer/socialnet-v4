package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

// @Summary Retrieve site details
// @Description This endpoint retrieves few details of the site/application.
// @Tags General
// @Success 200 {object} schemas.SiteDetailResponseSchema
// @Router /general/site-detail [get]
func (endpoint Endpoint) GetSiteDetails(c *fiber.Ctx) error {
	db := endpoint.DB
	sitedetail := managers.SiteDetailManager{}.GetOrCreate(db)
	data := utils.ConvertStructData(sitedetail, schemas.SiteDetail{}).(*schemas.SiteDetail)
	responseSiteDetail := schemas.SiteDetailResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Site Details Fetched!"}.Init(),
		Data:           *data,
	}
	return c.Status(200).JSON(responseSiteDetail)
}
