package routes

import (
	"math"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

func PaginateQueryset(queryset interface{}, c *fiber.Ctx, opts ...int) (schemas.PaginatedResponseDataSchema, interface{}, *string) {
	currentPage := c.QueryInt("page", 1)
	var perPage int
	var err *string = nil

	if currentPage < 1 {
		e := "Invalid Page"
		err = &e
	}

	// Check if page size is provided as an argument
	if len(opts) > 0 {
		perPage = opts[0]
	} else {
		// Default page size if not provided
		perPage = 50
	}
	querysetValue := reflect.ValueOf(queryset)
	itemsCount := querysetValue.Len()
	lastPage := math.Ceil(float64(itemsCount) / float64(perPage))
	if lastPage == 0 {
		lastPage = 1
	}

	startIndex := (currentPage - 1) * perPage
	endIndex := startIndex + perPage

	// Ensure startIndex is within bounds
	if startIndex < 0 {
		startIndex = 0
	}

	// Ensure endIndex is within bounds
	if endIndex > itemsCount {
		endIndex = itemsCount
	}

	paginatorData := schemas.PaginatedResponseDataSchema{
		PerPage:     uint(perPage),
		CurrentPage: uint(currentPage),
		LastPage:    uint(lastPage),
	}
	paginatedItems := querysetValue.Slice(startIndex, endIndex).Interface()
	
	if currentPage > int(lastPage) {
		e := "Page number is out of range"
		err = &e
	}
	return paginatorData, paginatedItems, err
}
