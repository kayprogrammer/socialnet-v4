package tests

import (
	"time"

	auth "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/friend"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

var (
	userManager = managers.UserManager{}
	friendManager = managers.FriendManager{}
)

func CreateTestUser(db *ent.Client) *ent.User {
	userData := schemas.RegisterUser{
		FirstName:      "Test",
		LastName:       "User",
		Email:          "testuser@example.com",
		Password:       "testpassword",
		TermsAgreement: false,
	}
	user := userManager.GetOrCreate(db, userData, false, false)
	return user
}

func CreateTestVerifiedUser(db *ent.Client) *ent.User {
	userData := schemas.RegisterUser{
		FirstName: "Test",
		LastName:  "Verified",
		Email:     "testverifieduser@example.com",
		Password:  "testpassword",
	}
	user := userManager.GetOrCreate(db, userData, true, false)
	return user
}

func CreateAnotherTestVerifiedUser(db *ent.Client) *ent.User {
	userData := schemas.RegisterUser{
		FirstName: "AnotherTest",
		LastName:  "UserVerified",
		Email:     "anothertestverifieduser@example.com",
		Password:  "testpassword",
	}
	user := userManager.GetOrCreate(db, userData, true, false)
	return user
}

func CreateJwt(db *ent.Client, user *ent.User) *ent.User {
	user = userManager.UpdateTokens(user, auth.GenerateAccessToken(user.ID, user.Username), auth.GenerateRefreshToken())
	return user
}

func AccessToken(db *ent.Client) string {
	user := CreateTestVerifiedUser(db)
	user = CreateJwt(db, user)
	return *user.Access
}

func AnotherAccessToken(db *ent.Client) string {
	user := CreateAnotherTestVerifiedUser(db)
	user = CreateJwt(db, user)
	return *user.Access
}

func CreateCity(db *ent.Client) *ent.City {
	country := managers.CountryManager{}.Create(db, "Nigeria", "NG")
	region := managers.RegionManager{}.Create(db, "Lagos", country)
	city := managers.CityManager{}.Create(db, "Lekki", country, region)
	return city
}

func CreateFriend(db *ent.Client, status friend.Status) *ent.Friend {
	verifiedUser := CreateTestVerifiedUser(db)
	anotherVerifiedUser := CreateAnotherTestVerifiedUser(db)
	friend := friendManager.Create(db, verifiedUser, anotherVerifiedUser, status)
	return friend
}

func ConvertDateTime (timeObj time.Time) string {
	return timeObj.Round(time.Microsecond).Format("2006-01-02T15:04:05.000000-07:00")
}