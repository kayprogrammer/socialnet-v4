package sockets

import (
	"encoding/json"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

type SocketNotificationSchema struct {
	schemas.NotificationSchema
	Status				string			`json:"status"`
}

// Maintain db & a list of connected clients
var (
	clients = make(map[*websocket.Conn]bool)
	clientsMutex = &sync.Mutex{}
	notification SocketNotificationSchema
)

// Function to add a client to the list
func AddClient(c *websocket.Conn) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()
    clients[c] = true
}

// Function to remove a client from the list
func RemoveClient(c *websocket.Conn) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()
    delete(clients, c)
}

type ErrorResp struct {
	Code			uint		`json:"code"`
	Message			string		`json:"message"`
}

func ReturnError(c *websocket.Conn, code uint, message string) {
	errorResponse := ErrorResp{Code: code, Message: message}
	jsonResponse, _ := json.Marshal(errorResponse)
	c.WriteMessage(websocket.TextMessage, jsonResponse)
}

func ValidateAuth (db *ent.Client, token string) (*ent.User, *string, *string) {
	var (
		errMsg *string
		secret *string
		user *ent.User
	)
	if len(token) < 1 {
		err := "Auth bearer not set"
		errMsg = &err
	} else if token == config.GetConfig().SocketSecret {
		secret = &token 
	} else {
		// Get User
		userObj, err := authentication.GetUser(token, db)
		if err != nil {
			errMsg = err
		}
		user = userObj
	}
	return user, secret, errMsg
}