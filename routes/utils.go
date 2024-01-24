package routes

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func ValidateReactionFocus(focus string) *utils.ErrorResponse {
	switch focus {
		case "POST", "COMMENT", "REPLY": return nil
	}
	err := utils.RequestErr(utils.ERR_INVALID_VALUE, "Invalid 'focus' value")
	return &err 
}

type SocketNotificationSchema struct {
	schemas.NotificationSchema
	Status				string			`json:"status"`
}

func SendNotificationInSocket(isSecured bool, host string, notification *ent.Notification, statusOpts ...string) error {
	// Check if page size is provided as an argument
	status := "CREATED"
	if len(statusOpts) > 0 {
		status = statusOpts[0]
	}
	webSocketScheme := "ws://"
	if isSecured {
		webSocketScheme = "wss://"
	}
	uri := webSocketScheme + host + "/api/v4/ws/notifications/"
	notificationData := SocketNotificationSchema{
		NotificationSchema: schemas.NotificationSchema{ID: notification.ID, Ntype: string(notification.Ntype)},
		Status: status,
	}
	if status == "CREATED" {
		convertedNotification := utils.ConvertStructData(notification, schemas.NotificationSchema{}).(*schemas.NotificationSchema)
		notificationData = SocketNotificationSchema{
			NotificationSchema: convertedNotification.Init(nil),
			Status: status,
		}
	}

	// Connect to the WebSocket server
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	headers := make(http.Header)
	headers.Add("Authorization", config.GetConfig().SocketSecret)
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