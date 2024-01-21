package schemas

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
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