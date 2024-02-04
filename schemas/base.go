package schemas

import (
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type ResponseSchema struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Data fetched/created/updated/deleted"`
}

func (obj ResponseSchema) Init() ResponseSchema {
	if obj.Status == "" {
		obj.Status = "success"
	}
	return obj
}

func FullName(user *ent.User) string {
	return user.FirstName + " " + user.LastName
}

func AvatarUrl(user *ent.User) *string {
	avatar := user.Edges.Avatar
	if avatar != nil {
		url := utils.GenerateFileUrl(avatar.ID.String(), "avatars", avatar.ResourceType)
		return &url
	}
	return nil
}

type PaginatedResponseDataSchema struct {
	PerPage     uint `json:"per_page" example:"100"`
	CurrentPage uint `json:"current_page" example:"1"`
	LastPage    uint `json:"last_page" example:"100"`
}

type UserDataSchema struct {
	Name     string  `json:"name" example:"John Doe"`
	Username string  `json:"username" example:"john-doe"`
	Avatar   *string `json:"avatar" example:"https://img.url"`
}

func (user UserDataSchema) Init(userObj *ent.User) UserDataSchema {
	user.Name = FullName(userObj)
	user.Username = userObj.Username
	user.Avatar = AvatarUrl(userObj)
	return user
}
