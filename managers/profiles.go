package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/city"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
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

func (obj UserProfileManager) GetByUsername(client *ent.Client, username string) (*ent.User, *utils.ErrorResponse) {
	u, err := client.User.
		Query().
		Where(user.Username(username)).
		WithAvatar().
		WithCity().
		Only(Ctx)
	if err != nil {
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "No user with that username")
		return nil, &errData
	}
	return u, nil
}

func (obj UserProfileManager) Update(client *ent.Client, profile *ent.User, profileData schemas.ProfileUpdateSchema) *ent.User {
	var avatarId *uuid.UUID
	avatar := profile.Edges.Avatar
	fileM := FileManager{}
	if profileData.FileType != nil {
		// Create or Update Image Object
		if avatar == nil {
			avatar, _ = FileManager{}.Create(client, profileData.FileType)
		} else {
			avatar = fileM.Update(client, avatar, *profileData.FileType)

		}
		avatarId = &avatar.ID
	}

	u, _ := profile.
		Update().
		SetNillableFirstName(profileData.FirstName).
		SetNillableLastName(profileData.LastName).
		SetNillableBio(profileData.Bio).
		SetNillableDob(profileData.Dob).
		SetNillableCityID(profileData.CityID).
		SetNillableAvatarID(avatarId).
		Save(Ctx)

	// Set related values
	city := profileData.City
	if city == nil {
		city = profile.Edges.City
	}
	u.Edges.City = city
	u.Edges.Avatar = avatar
	return u
}
