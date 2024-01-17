package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
)

// Validates if a date has a correct format (ISO8601)
func DateValidator(fl validator.FieldLevel) bool {
	inputTimeString := fl.Field().String()
	_, err := time.Parse(time.RFC3339, inputTimeString)
	return err == nil
}

// Validates if a reaction value is the correct one
func ReactionTypeValidator(fl validator.FieldLevel) bool {
	reactionValue := fl.Field().Interface().(reaction.Rtype)
	switch reactionValue {
	case reaction.RtypeLIKE, reaction.RtypeLOVE, reaction.RtypeHAHA, reaction.RtypeWOW, reaction.RtypeSAD, reaction.RtypeANGRY:
		return true
	}
	return false // Error. Value doesn't match the required
}

// Validates if a file type is accepted
func FileTypeValidator(fl validator.FieldLevel) bool {
	fileType := fl.Field().Interface().(string)
	fileTypeFound := false
	for key := range ImageExtensions {
		if key == fileType {
			fileTypeFound = true
			break
		}
	}
	return fileTypeFound
}
