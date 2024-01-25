package sockets

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/kayprogrammer/socialnet-v4/database"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/comment"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

type SocketNotificationSchema struct {
	schemas.NotificationSchema
	Status string `json:"status"`
}

var notificationObj SocketNotificationSchema

// Function to broadcast a notification data to all connected clients
func broadcastNotificationMessage(db *ent.Client, mt int, msg []byte) {
	notificationManager := managers.NotificationManager{}

	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		user := client.Locals("user").(*ent.User)
		if user == nil {
			continue
		}
		json.Unmarshal(msg, &notificationObj)
		// Ensure user is a valid recipient of this notification
		userIsAmongReceiver := notificationManager.IsAmongReceivers(db, notificationObj.ID, user.ID)
		if userIsAmongReceiver {
			if err := client.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
			}
		}
	}
	// Delete comment or reply here after the socket message has been sent for comment & reply deletion
	// Although another better way will be to delete the comment or reply the respective view/handler
	// But then the notification will be deleted alongside (cos of CASCADE relationship) before the notification socket will be sent
	// Which will prevent the user from seeing the real time notification cos the IsAmongReceivers won't work with an already deleted notifiation
	// To prevent this you can just set the relationship to SetNull, then delete notification here, and delete comment & reply in the view.
	// The only drawback I can think of concerning the below method is that if by any means there was an issue with the socket, the stuff won't get deleted.
	// Omo na wahala be that oh. But anyway, just go ahead with the SetNull whatever. I'm too lazy to change anything now.
	// Sorry for the long note (no vex)
	if notificationObj.Status == "DELETED" && notificationObj.Ntype != "REACTION" {
		if notificationObj.CommentSlug != nil {
			db.Comment.Delete().Where(comment.Slug(*notificationObj.CommentSlug)).ExecX(managers.Ctx)
		} else if notificationObj.ReplySlug != nil {
			db.Reply.Delete().Where(reply.Slug(*notificationObj.ReplySlug)).ExecX(managers.Ctx)
		}
	}
}

func NotificationSocket(c *websocket.Conn) {
	db := database.ConnectDb()
	defer db.Close()
	token := c.Headers("Authorization")

	var (
		mt     int
		msg    []byte
		err    error
		user   *ent.User
		secret *string
		errM   *string
	)

	// Validate Auth
	if user, secret, errM = ValidateAuth(db, token); errM != nil {
		ReturnError(c, 4001, *errM)
		return
	}
	// Add the client to the list of connected clients
	c.Locals("user", user)
	AddClient(c)

	// Remove the client from the list when the handler exits
	defer RemoveClient(c)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		// Notifications can only be broadcasted from the app using the socket secret
		if secret != nil {
			broadcastNotificationMessage(db, mt, msg)
		} else {
			ReturnError(c, 4001, "Not authorized to send data")
			break
		}
	}
}
