package managers

import (
	"github.com/gosimple/slug"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func UsernameGenerator(db *ent.Client, firstName string, lastName string, userName *string) string {
	uniqueUsername := slug.Make(firstName + " " + lastName)
	if userName != nil {
		uniqueUsername = *userName
	}

	obj, _ := db.User.Query().Where(user.Username(uniqueUsername)).Only(Ctx)
	if obj != nil { // username is already taken
		// Make it unique by attaching a random string
		// to it and repeat the function
		randomStr := utils.GetRandomString(6)
		uniqueUsername = uniqueUsername + "-" + randomStr
		return UsernameGenerator(db, firstName, lastName, &uniqueUsername)
	}
	return uniqueUsername
}
