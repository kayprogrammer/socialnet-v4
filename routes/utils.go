package routes

import "github.com/kayprogrammer/socialnet-v4/utils"

func ValidateReactionFocus(focus string) *utils.ErrorResponse {
	expectedFocuses := []string{"POST", "COMMENT", "REPLY"}
	found := false
	for _, str := range expectedFocuses {
		if str == focus {
			found = true
			break
		}
	}
	if !found {
		err := utils.RequestErr(utils.ERR_INVALID_VALUE, "Invalid 'focus' value")
		return &err 
	}
	return nil
}