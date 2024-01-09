package managers

import (
	"strings"

	"github.com/gosimple/slug"
	"github.com/kayprogrammer/socialnet-v4/managers"
)

func UsernameGenerator(firstName string, lastName string, userName *string) {
	name := firstName + " " + lastName
	if name && (!username || !strings.HasPrefix(username, slug.Make(name))) {
		// The if statement above implies that username will only be created or altered
		// if name exists and
		// if username is none OR name has changed (checking if the current username tallies with the name)

		uniqueUsername = slug.Make(name)
		obj := managers.UserManager{}.GetByIDAndUsername(db, id, uniqueUsername)
		obj = (
			await User.objects()
			.where(User.username == unique_username, User.id != self.id)
			.first()
		)
		if obj:
			unique_username = (
				f"{unique_username}-{generate_random_alphanumeric_string()}"
			)
			return await self.generate_username(unique_username)
		return unique_username
	}
	return username 
}