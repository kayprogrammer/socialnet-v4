package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type CitySchema struct {
	Edges        	*ent.CityEdges 	`json:"edges,omitempty" swaggerignore:"true"`
    ID 				uuid.UUID			`json:"id" example:"d10dde64-a242-4ed0-bd75-4c759644b3a6"`
    Name 			string				`json:"name" example:"Lekki"`
    Region 			*string				`json:"region" example:"Lagos"`
    Country 		string				`json:"country" example:"Nigeria"`
}

func (city CitySchema) Init () CitySchema {
	// Set Related Data.
	region := city.Edges.Region
	if region != nil {
		city.Region = &region.Name
	}
	city.Country = city.Edges.Country.Name

	city.Edges = nil // Omit edges
	return city
}

type ProfileSchema struct {
	Edges        	*ent.UserEdges 		`json:"edges,omitempty" swaggerignore:"true"`
    FirstName 		string				`json:"first_name" example:"John"`
    LastName 		string				`json:"last_name" example:"Doe"`
    Username 		string				`json:"username" example:"john-doe"`
    Email 			string				`json:"email" example:"johndoe@email.com"`
    Avatar 			*string				`json:"avatar" example:"https://img.com"`
    Bio 			*string				`json:"bio" example:"Software Engineer | Django Ninja Developer"`
    Dob 			*time.Time			`json:"dob"`
    City 			*string				`json:"city" example:"Lekki"`
    CreatedAt 		*time.Time			`json:"created_at"`
    UpdatedAt 		*time.Time			`json:"updated_at"`
}

func (user ProfileSchema) Init () ProfileSchema {
	// Set Related Data.
	avatar := user.Edges.Avatar 
	if avatar != nil {
		url := utils.GenerateFileUrl(avatar.ID.String(), "avatars", avatar.ResourceType)
		user.Avatar = &url
	}
	city := user.Edges.City
	if city != nil {
		user.City = &city.Name
	}
	user.Edges = nil // Omit edges
	return user
}

// RESPONSE SCHEMAS
// CITIES
type CitiesResponseSchema struct {
	ResponseSchema
	Data			[]CitySchema		`json:"data"`
}


func (data CitiesResponseSchema) Init () CitiesResponseSchema {
	// Set Initial Data
	cities := data.Data
	for i := range cities {
		cities[i] = cities[i].Init()
	}
	data.Data = cities
	return data
}

// USERS
type ProfilesResponseDataSchema struct {
	PaginatedResponseDataSchema
	Items			[]ProfileSchema		`json:"users"`
}

func (data ProfilesResponseDataSchema) Init () ProfilesResponseDataSchema {
	// Set Initial Data
	items := data.Items
	for i := range items {
		items[i] = items[i].Init()
	}
	data.Items = items
	return data
}

type ProfilesResponseSchema struct {
	ResponseSchema
	Data			ProfilesResponseDataSchema		`json:"data"`
}
