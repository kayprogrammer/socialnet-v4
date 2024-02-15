package routes

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/sockets"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var cfg = config.GetConfig()

func ValidateReactionFocus(focus string) *utils.ErrorResponse {
	switch focus {
	case "POST", "COMMENT", "REPLY":
		return nil
	}
	err := utils.RequestErr(utils.ERR_INVALID_VALUE, "Invalid 'focus' value")
	return &err
}

func SendNotificationInSocket(fiberCtx *fiber.Ctx, notification *ent.Notification, commentSlug *string, replySlug *string, statusOpts ...string) error {
	if os.Getenv("ENVIRONMENT") == "TESTING" {
		return nil
	}
	
	// Check if page size is provided as an argument
	status := "CREATED"
	if len(statusOpts) > 0 {
		status = statusOpts[0]
	}
	webSocketScheme := "ws://"
	if fiberCtx.Secure() {
		webSocketScheme = "wss://"
	}
	uri := webSocketScheme + fiberCtx.Hostname() + "/api/v4/ws/notifications/"
	notificationData := sockets.SocketNotificationSchema{
		NotificationSchema: schemas.NotificationSchema{ID: notification.ID, Ntype: string(notification.Ntype), CommentSlug: commentSlug, ReplySlug: replySlug},
		Status:             status,
	}
	if status == "CREATED" {
		convertedNotification := utils.ConvertStructData(notification, schemas.NotificationSchema{}).(*schemas.NotificationSchema)
		notificationData = sockets.SocketNotificationSchema{
			NotificationSchema: convertedNotification.Init(nil),
			Status:             status,
		}
	}

	// Connect to the WebSocket server
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	headers := make(http.Header)
	headers.Add("Authorization", cfg.SocketSecret)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Marshal the notification data to JSON
	data, err := json.Marshal(notificationData)
	if err != nil {
		return err
	}

	// Send the notification to the WebSocket server
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}

	// Close the WebSocket connection
	return conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func SendMessageDeletionInSocket(fiberCtx *fiber.Ctx, chatID uuid.UUID, messageID uuid.UUID) error {
	if os.Getenv("ENVIRONMENT") == "TESTING" {
		return nil
	}
	webSocketScheme := "ws://"
	if fiberCtx.Secure() {
		webSocketScheme = "wss://"
	}
	uri := webSocketScheme + fiberCtx.Hostname() + "/api/v4/ws/chats/" + chatID.String()
	chatData := sockets.SocketMessageEntrySchema{
		ID:     messageID,
		Status: "DELETED",
	}

	// Connect to the WebSocket server
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	headers := make(http.Header)
	headers.Add("Authorization", cfg.SocketSecret)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Marshal the notification data to JSON
	data, err := json.Marshal(chatData)
	if err != nil {
		return err
	}

	// Send the message to the WebSocket server
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}

	// Close the WebSocket connection
	return conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
