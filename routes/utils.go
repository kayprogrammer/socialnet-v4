package routes

import (
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func GetPostObject(client *ent.Client, slug string, detailed bool) (*ent.Post, *int, *utils.ErrorResponse) {
	q := client.Post.Query().Where()
	if detailed {
		q = q.WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithImage().
			WithReactions().
			WithComments()
	}
	post, _ := q.Only(managers.Ctx)
	if post == nil {
		status_code := 404
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "Post does not exist")
		return nil, &status_code, &errData
	}
	return post, nil, nil
}