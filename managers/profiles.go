package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/city"
)

// ----------------------------------
// CITY MANAGEMENT
// --------------------------------
type CityManager struct {
}

func (obj CityManager) All(client *ent.Client, name string) []*ent.City {
	cities, _ := client.City.Query().
		Where(city.NameContains(name)).
		WithCountry().
		WithRegion().
		Limit(10).
		All(Ctx)
	return cities
}

func (obj CityManager) GetByID(client *ent.Client, cityID uuid.UUID) *ent.City {
	c, _ := client.City.Query().
		Where(city.ID(cityID)).
		Only(Ctx)
	return c
}