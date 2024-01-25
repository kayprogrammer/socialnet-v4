package sockets

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/kayprogrammer/socialnet-v4/database"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
)

// Function to broadcast a notification data to all connected clients
func broadcastNotificationMessage(db *ent.Client, mt int, msg []byte) {
	notificationManager := managers.NotificationManager{}
    clientsMutex.Lock()
    defer clientsMutex.Unlock()
    for client := range clients {
		user := client.Locals("user").(*ent.User)
		json.Unmarshal(msg, &notification)
		// Ensure user is a valid recipient of this notification
		if user != nil && notificationManager.IsAmongReceivers(db, notification.ID, user.ID){
			if err := client.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
			}
		}
    }
}

func NotificationSocket (c *websocket.Conn) {
	db := database.ConnectDb()
	defer db.Close()
	token := c.Headers("Authorization")

	var (
		mt  int
		msg []byte
		err error
		user *ent.User
		secret *string
		errM *string
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

