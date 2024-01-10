package managers

import (
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func UsernameGenerator(db *ent.Client, firstName string, lastName string, userId *uuid.UUID, userName *string) string {
	var uniqueUsername string
	name := firstName + " " + lastName
	if userName == nil || !(strings.HasPrefix(*userName, slug.Make(name))) {
		// The if statement above implies that username will only be created or altered
		// if username is none OR name has changed (checking if the current username tallies with the name)
		uniqueUsername = slug.Make(name)
		if userName != nil {
			uniqueUsername = *userName
		}

		// Exclude the current user id during the query if it exists
		obj, _ := db.User.Query().Where(user.Username(uniqueUsername)).Only(Ctx)
		if userId != nil {
			obj, _ = db.User.Query().Where(user.Not(user.ID(*userId))).Only(Ctx)
		}

		if obj != nil {
			// If there's another row with the slug, attach a random string to it 
			// to ensure its uniqueness and repeat the function
			randomStr := utils.GetRandomString(6)
			uniqueUsername = uniqueUsername + "-" + randomStr
			return UsernameGenerator(db, firstName, lastName, userId, &uniqueUsername)
		}
		return uniqueUsername
	}
	return *userName
}
