package initials

import (
	"log"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)
// var	truth = true
var userManager = managers.UserManager{}

func createSuperUser(db *ent.Client) *ent.User {
	userData := schemas.RegisterUser{
		FirstName:       "Test",
		LastName:        "Admin",
		Email:           "testadmin@email.com",
		Password:        "testadmin",
		TermsAgreement:  true,
	}
	user := userManager.GetOrCreate(db, userData, true, true)
	return user
}

func createClient(db *ent.Client) *ent.User {
	userData := schemas.RegisterUser{
		FirstName:       "Test",
		LastName:        "Client",
		Email:           "testclient@email.com",
		Password:        "testclient",
		TermsAgreement:  true,
	}
	user := userManager.GetOrCreate(db, userData, true, false)
	return user
}

func createSiteDetail(db *ent.Client) {
	managers.SiteDetailManager{}.GetOrCreate(db)
}


func CreateInitialData(db *ent.Client) {
	log.Println("Creating Initial Data....")
	createSuperUser(db)
	createClient(db)
	createSiteDetail(db)
	log.Println("Initial Data Created....")
}