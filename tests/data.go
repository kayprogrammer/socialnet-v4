package tests

import (
	auth "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

var (
	userManager = managers.UserManager{}
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
