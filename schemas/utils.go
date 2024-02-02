package schemas

import (
	"github.com/kayprogrammer/socialnet-v4/ent"
)

func ConvertUsers(users []*ent.User) []*UserDataSchema {
	convertedUsers := []*UserDataSchema{}
	for i := range users {
		user := UserDataSchema{}.Init(users[i])
		convertedUsers = append(convertedUsers, &user)
	}
	return convertedUsers
}