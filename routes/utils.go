package routes

import (
	"github.com/kayprogrammer/socialnet-v4/utils"
)

func ValidateReactionFocus(focus string) *utils.ErrorResponse {
	switch focus {
		case "POST", "COMMENT", "REPLY": return nil
	}
	err := utils.RequestErr(utils.ERR_INVALID_VALUE, "Invalid 'focus' value")
	return &err 
}