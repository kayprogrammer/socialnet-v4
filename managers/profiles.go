package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/city"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
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

// ----------------------------------
// USER PROFILE MANAGEMENT
// --------------------------------
type UserProfileManager struct {
}

func (obj UserProfileManager) GetUsers(client *ent.Client, userObj *ent.User) []*ent.User {
	uq := client.User.Query().
		WithAvatar().
		WithCity()
	if userObj != nil {
		// Exclude yourself
		uq = uq.Where(user.Not(user.ID(userObj.ID)))
	}
	users, _ := uq.All(Ctx)
	return users
}
