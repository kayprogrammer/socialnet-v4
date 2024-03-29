package tests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	auth "github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/friend"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

var (
	userManager         = managers.UserManager{}
	friendManager       = managers.FriendManager{}
	notificationManager = managers.NotificationManager{}
	chatManager         = managers.ChatManager{}
	messageManager      = managers.MessageManager{}
	postManager         = managers.PostManager{}
	reactionManager     = managers.ReactionManager{}
	commentManager      = managers.CommentManager{}
	replyManager        = managers.ReplyManager{}
)

// AUTH FIXTURES
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

// ----------------------------------------------------------------------------

// PROFILE FIXTURES
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

func CreateNotification(db *ent.Client) *ent.Notification {
	user := CreateTestVerifiedUser(db)
	text := "A new update is coming!"
	notification := notificationManager.Create(db, nil, "ADMIN", []uuid.UUID{user.ID}, nil, nil, nil, &text)
	return notification
}

// ----------------------------------------------------------------------------

// CHAT FIXTURES
func CreateChat(db *ent.Client) *ent.Chat {
	verifiedUser := CreateTestVerifiedUser(db)
	anotherVerifiedUser := CreateAnotherTestVerifiedUser(db)
	chat := chatManager.GetDMChat(db, verifiedUser, anotherVerifiedUser)
	if chat == nil {
		chat = chatManager.Create(db, verifiedUser, "DM", []*ent.User{anotherVerifiedUser})
	} else {
		// Set useful related data
		chat.Edges.Owner = verifiedUser
	}
	chat.Edges.Users = []*ent.User{anotherVerifiedUser}
	return chat
}

func CreateGroupChat(db *ent.Client) *ent.Chat {
	verifiedUser := CreateTestVerifiedUser(db)
	anotherVerifiedUser := CreateAnotherTestVerifiedUser(db)
	chatManager.DropData(db)
	dataToCreate := schemas.GroupChatCreateSchema{Name: "My New Group"}
	chat := chatManager.CreateGroup(db, verifiedUser, []*ent.User{anotherVerifiedUser}, dataToCreate)
	chat.Edges.Users = []*ent.User{anotherVerifiedUser}
	return chat
}

func CreateMessage(db *ent.Client) *ent.Message {
	messageManager.DropData(db)
	chat := CreateChat(db)
	text := "Hello Boss"
	message := messageManager.Create(db, chat.Edges.Owner, chat, &text, nil)
	return message
}

// ----------------------------------------------------------------------------

// FEED FIXTURES
func CreatePost(db *ent.Client) *ent.Post {
	author := CreateTestVerifiedUser(db)
	post := postManager.Create(db, author, schemas.PostInputSchema{Text: "This is a nice new platform."})
	return post
}

func CreateReaction(db *ent.Client) *ent.Reaction {
	post := CreatePost(db)
	reaction := reactionManager.Create(db, post.AuthorID, "POST", post.ID, "LIKE")
	reaction.Edges.Post = post
	reaction.Edges.User = post.Edges.Author
	return reaction
}

func CreateComment(db *ent.Client) *ent.Comment {
	post := CreatePost(db)
	comment := commentManager.Create(db, post.Edges.Author, post.ID, "Just a comment")
	comment.Edges.Post = post
	return comment
}

func CreateReply(db *ent.Client) *ent.Reply {
	comment := CreateComment(db)
	reply := replyManager.Create(db, comment.Edges.Author, comment.ID, "Simple reply")
	reply.Edges.Comment = comment
	return reply
}

// ----------------------------------------------------------------------------

// Utils
func GetUserMap(user *ent.User) map[string]interface{} {
	return map[string]interface{}{
		"name":     schemas.FullName(user),
		"username": user.Username,
		"avatar":   nil,
	}
}

func ConvertDateTime(timeObj time.Time) string {
	roundedTime := timeObj.Round(time.Microsecond)
	formatted := roundedTime.Format("2006-01-02T15:04:05")

	// Get the microsecond part and round it
	microseconds := roundedTime.Nanosecond() / 1000

	// Append the rounded microsecond part to the formatted string
	formatted = fmt.Sprintf("%s.%06d", formatted, microseconds)
	formatted = strings.TrimRight(formatted, "0")
	// Append the timezone information
	formatted = fmt.Sprintf("%s%s", formatted, roundedTime.Format("-07:00"))

	return formatted
}

// ----------------------------------------------------------------------------
