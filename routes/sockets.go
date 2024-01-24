package routes

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/kayprogrammer/socialnet-v4/authentication"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
)

type ErrorResp struct {
	Code			uint		`json:"code"`
	Message			string		`json:"message"`
}

func ValidateAuth (db *ent.Client, token string) (*ent.User, *string, *string) {
	cfg := config.GetConfig()
	var (
		errMsg *string
		secret *string
		user *ent.User
	)
	if len(token) < 1 {
		err := "Auth bearer not set"
		errMsg = &err
	} else if token == cfg.SocketSecret {
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

func ReturnError(c *websocket.Conn, code uint, message string) {
	errorResponse := ErrorResp{Code: code, Message: message}
	jsonResponse, _ := json.Marshal(errorResponse)
	c.WriteMessage(websocket.TextMessage, jsonResponse)
}

func NotificationSocket (c *websocket.Conn) {
	db := c.Locals("db").(*ent.Client)
	token := c.Headers("Authorization")
	// user, secret, errM := ValidateAuth(db, token) 

	var (
		mt  int
		msg []byte
		err error
		user *ent.User
		secret *string
		errM *string
	)
	for {
		// Validate Auth
		if user, secret, errM = ValidateAuth(db, token); errM != nil {
			log.Println("read:", )
			ReturnError(c, 4001, *errM)
			break
		}
		log.Printf("name: %s", user.FirstName)
		log.Printf("secret: %s", secret)

		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}

