// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/chat"
	"github.com/kayprogrammer/socialnet-v4/ent/city"
	"github.com/kayprogrammer/socialnet-v4/ent/comment"
	"github.com/kayprogrammer/socialnet-v4/ent/country"
	"github.com/kayprogrammer/socialnet-v4/ent/file"
	"github.com/kayprogrammer/socialnet-v4/ent/friend"
	"github.com/kayprogrammer/socialnet-v4/ent/message"
	"github.com/kayprogrammer/socialnet-v4/ent/notification"
	"github.com/kayprogrammer/socialnet-v4/ent/otp"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
	"github.com/kayprogrammer/socialnet-v4/ent/region"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/ent/schema"
	"github.com/kayprogrammer/socialnet-v4/ent/sitedetail"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescCreatedAt is the schema descriptor for created_at field.
	chatDescCreatedAt := chatFields[1].Descriptor()
	// chat.DefaultCreatedAt holds the default value on creation for the created_at field.
	chat.DefaultCreatedAt = chatDescCreatedAt.Default.(func() time.Time)
	// chatDescUpdatedAt is the schema descriptor for updated_at field.
	chatDescUpdatedAt := chatFields[2].Descriptor()
	// chat.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chat.DefaultUpdatedAt = chatDescUpdatedAt.Default.(func() time.Time)
	// chat.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chat.UpdateDefaultUpdatedAt = chatDescUpdatedAt.UpdateDefault.(func() time.Time)
	// chatDescID is the schema descriptor for id field.
	chatDescID := chatFields[0].Descriptor()
	// chat.DefaultID holds the default value on creation for the id field.
	chat.DefaultID = chatDescID.Default.(func() uuid.UUID)
	cityFields := schema.City{}.Fields()
	_ = cityFields
	// cityDescCreatedAt is the schema descriptor for created_at field.
	cityDescCreatedAt := cityFields[1].Descriptor()
	// city.DefaultCreatedAt holds the default value on creation for the created_at field.
	city.DefaultCreatedAt = cityDescCreatedAt.Default.(func() time.Time)
	// cityDescUpdatedAt is the schema descriptor for updated_at field.
	cityDescUpdatedAt := cityFields[2].Descriptor()
	// city.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	city.DefaultUpdatedAt = cityDescUpdatedAt.Default.(func() time.Time)
	// city.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	city.UpdateDefaultUpdatedAt = cityDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cityDescName is the schema descriptor for name field.
	cityDescName := cityFields[3].Descriptor()
	// city.NameValidator is a validator for the "name" field. It is called by the builders before save.
	city.NameValidator = cityDescName.Validators[0].(func(string) error)
	// cityDescID is the schema descriptor for id field.
	cityDescID := cityFields[0].Descriptor()
	// city.DefaultID holds the default value on creation for the id field.
	city.DefaultID = cityDescID.Default.(func() uuid.UUID)
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[1].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[2].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescText is the schema descriptor for text field.
	commentDescText := commentFields[3].Descriptor()
	// comment.TextValidator is a validator for the "text" field. It is called by the builders before save.
	comment.TextValidator = commentDescText.Validators[0].(func(string) error)
	// commentDescSlug is the schema descriptor for slug field.
	commentDescSlug := commentFields[4].Descriptor()
	// comment.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	comment.SlugValidator = commentDescSlug.Validators[0].(func(string) error)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.DefaultID holds the default value on creation for the id field.
	comment.DefaultID = commentDescID.Default.(func() uuid.UUID)
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescCreatedAt is the schema descriptor for created_at field.
	countryDescCreatedAt := countryFields[1].Descriptor()
	// country.DefaultCreatedAt holds the default value on creation for the created_at field.
	country.DefaultCreatedAt = countryDescCreatedAt.Default.(func() time.Time)
	// countryDescUpdatedAt is the schema descriptor for updated_at field.
	countryDescUpdatedAt := countryFields[2].Descriptor()
	// country.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	country.DefaultUpdatedAt = countryDescUpdatedAt.Default.(func() time.Time)
	// country.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	country.UpdateDefaultUpdatedAt = countryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// countryDescName is the schema descriptor for name field.
	countryDescName := countryFields[3].Descriptor()
	// country.NameValidator is a validator for the "name" field. It is called by the builders before save.
	country.NameValidator = countryDescName.Validators[0].(func(string) error)
	// countryDescCode is the schema descriptor for code field.
	countryDescCode := countryFields[4].Descriptor()
	// country.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	country.CodeValidator = countryDescCode.Validators[0].(func(string) error)
	// countryDescID is the schema descriptor for id field.
	countryDescID := countryFields[0].Descriptor()
	// country.DefaultID holds the default value on creation for the id field.
	country.DefaultID = countryDescID.Default.(func() uuid.UUID)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[1].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileFields[2].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescResourceType is the schema descriptor for resource_type field.
	fileDescResourceType := fileFields[3].Descriptor()
	// file.ResourceTypeValidator is a validator for the "resource_type" field. It is called by the builders before save.
	file.ResourceTypeValidator = fileDescResourceType.Validators[0].(func(string) error)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileFields[0].Descriptor()
	// file.DefaultID holds the default value on creation for the id field.
	file.DefaultID = fileDescID.Default.(func() uuid.UUID)
	friendFields := schema.Friend{}.Fields()
	_ = friendFields
	// friendDescCreatedAt is the schema descriptor for created_at field.
	friendDescCreatedAt := friendFields[1].Descriptor()
	// friend.DefaultCreatedAt holds the default value on creation for the created_at field.
	friend.DefaultCreatedAt = friendDescCreatedAt.Default.(func() time.Time)
	// friendDescUpdatedAt is the schema descriptor for updated_at field.
	friendDescUpdatedAt := friendFields[2].Descriptor()
	// friend.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	friend.DefaultUpdatedAt = friendDescUpdatedAt.Default.(func() time.Time)
	// friend.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	friend.UpdateDefaultUpdatedAt = friendDescUpdatedAt.UpdateDefault.(func() time.Time)
	// friendDescID is the schema descriptor for id field.
	friendDescID := friendFields[0].Descriptor()
	// friend.DefaultID holds the default value on creation for the id field.
	friend.DefaultID = friendDescID.Default.(func() uuid.UUID)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[1].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageFields[2].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[1].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationFields[2].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescID is the schema descriptor for id field.
	notificationDescID := notificationFields[0].Descriptor()
	// notification.DefaultID holds the default value on creation for the id field.
	notification.DefaultID = notificationDescID.Default.(func() uuid.UUID)
	otpFields := schema.Otp{}.Fields()
	_ = otpFields
	// otpDescCreatedAt is the schema descriptor for created_at field.
	otpDescCreatedAt := otpFields[1].Descriptor()
	// otp.DefaultCreatedAt holds the default value on creation for the created_at field.
	otp.DefaultCreatedAt = otpDescCreatedAt.Default.(func() time.Time)
	// otpDescUpdatedAt is the schema descriptor for updated_at field.
	otpDescUpdatedAt := otpFields[2].Descriptor()
	// otp.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	otp.DefaultUpdatedAt = otpDescUpdatedAt.Default.(func() time.Time)
	// otp.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	otp.UpdateDefaultUpdatedAt = otpDescUpdatedAt.UpdateDefault.(func() time.Time)
	// otpDescID is the schema descriptor for id field.
	otpDescID := otpFields[0].Descriptor()
	// otp.DefaultID holds the default value on creation for the id field.
	otp.DefaultID = otpDescID.Default.(func() uuid.UUID)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[1].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[2].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescText is the schema descriptor for text field.
	postDescText := postFields[3].Descriptor()
	// post.TextValidator is a validator for the "text" field. It is called by the builders before save.
	post.TextValidator = postDescText.Validators[0].(func(string) error)
	// postDescSlug is the schema descriptor for slug field.
	postDescSlug := postFields[4].Descriptor()
	// post.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	post.SlugValidator = postDescSlug.Validators[0].(func(string) error)
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.DefaultID holds the default value on creation for the id field.
	post.DefaultID = postDescID.Default.(func() uuid.UUID)
	reactionFields := schema.Reaction{}.Fields()
	_ = reactionFields
	// reactionDescCreatedAt is the schema descriptor for created_at field.
	reactionDescCreatedAt := reactionFields[1].Descriptor()
	// reaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	reaction.DefaultCreatedAt = reactionDescCreatedAt.Default.(func() time.Time)
	// reactionDescUpdatedAt is the schema descriptor for updated_at field.
	reactionDescUpdatedAt := reactionFields[2].Descriptor()
	// reaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reaction.DefaultUpdatedAt = reactionDescUpdatedAt.Default.(func() time.Time)
	// reaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reaction.UpdateDefaultUpdatedAt = reactionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// reactionDescID is the schema descriptor for id field.
	reactionDescID := reactionFields[0].Descriptor()
	// reaction.DefaultID holds the default value on creation for the id field.
	reaction.DefaultID = reactionDescID.Default.(func() uuid.UUID)
	regionFields := schema.Region{}.Fields()
	_ = regionFields
	// regionDescCreatedAt is the schema descriptor for created_at field.
	regionDescCreatedAt := regionFields[1].Descriptor()
	// region.DefaultCreatedAt holds the default value on creation for the created_at field.
	region.DefaultCreatedAt = regionDescCreatedAt.Default.(func() time.Time)
	// regionDescUpdatedAt is the schema descriptor for updated_at field.
	regionDescUpdatedAt := regionFields[2].Descriptor()
	// region.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	region.DefaultUpdatedAt = regionDescUpdatedAt.Default.(func() time.Time)
	// region.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	region.UpdateDefaultUpdatedAt = regionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// regionDescName is the schema descriptor for name field.
	regionDescName := regionFields[3].Descriptor()
	// region.NameValidator is a validator for the "name" field. It is called by the builders before save.
	region.NameValidator = regionDescName.Validators[0].(func(string) error)
	// regionDescID is the schema descriptor for id field.
	regionDescID := regionFields[0].Descriptor()
	// region.DefaultID holds the default value on creation for the id field.
	region.DefaultID = regionDescID.Default.(func() uuid.UUID)
	replyFields := schema.Reply{}.Fields()
	_ = replyFields
	// replyDescCreatedAt is the schema descriptor for created_at field.
	replyDescCreatedAt := replyFields[1].Descriptor()
	// reply.DefaultCreatedAt holds the default value on creation for the created_at field.
	reply.DefaultCreatedAt = replyDescCreatedAt.Default.(func() time.Time)
	// replyDescUpdatedAt is the schema descriptor for updated_at field.
	replyDescUpdatedAt := replyFields[2].Descriptor()
	// reply.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reply.DefaultUpdatedAt = replyDescUpdatedAt.Default.(func() time.Time)
	// reply.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reply.UpdateDefaultUpdatedAt = replyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// replyDescText is the schema descriptor for text field.
	replyDescText := replyFields[3].Descriptor()
	// reply.TextValidator is a validator for the "text" field. It is called by the builders before save.
	reply.TextValidator = replyDescText.Validators[0].(func(string) error)
	// replyDescSlug is the schema descriptor for slug field.
	replyDescSlug := replyFields[4].Descriptor()
	// reply.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	reply.SlugValidator = replyDescSlug.Validators[0].(func(string) error)
	// replyDescID is the schema descriptor for id field.
	replyDescID := replyFields[0].Descriptor()
	// reply.DefaultID holds the default value on creation for the id field.
	reply.DefaultID = replyDescID.Default.(func() uuid.UUID)
	sitedetailFields := schema.SiteDetail{}.Fields()
	_ = sitedetailFields
	// sitedetailDescCreatedAt is the schema descriptor for created_at field.
	sitedetailDescCreatedAt := sitedetailFields[1].Descriptor()
	// sitedetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	sitedetail.DefaultCreatedAt = sitedetailDescCreatedAt.Default.(func() time.Time)
	// sitedetailDescUpdatedAt is the schema descriptor for updated_at field.
	sitedetailDescUpdatedAt := sitedetailFields[2].Descriptor()
	// sitedetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sitedetail.DefaultUpdatedAt = sitedetailDescUpdatedAt.Default.(func() time.Time)
	// sitedetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sitedetail.UpdateDefaultUpdatedAt = sitedetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sitedetailDescName is the schema descriptor for name field.
	sitedetailDescName := sitedetailFields[3].Descriptor()
	// sitedetail.DefaultName holds the default value on creation for the name field.
	sitedetail.DefaultName = sitedetailDescName.Default.(string)
	// sitedetailDescEmail is the schema descriptor for email field.
	sitedetailDescEmail := sitedetailFields[4].Descriptor()
	// sitedetail.DefaultEmail holds the default value on creation for the email field.
	sitedetail.DefaultEmail = sitedetailDescEmail.Default.(string)
	// sitedetailDescPhone is the schema descriptor for phone field.
	sitedetailDescPhone := sitedetailFields[5].Descriptor()
	// sitedetail.DefaultPhone holds the default value on creation for the phone field.
	sitedetail.DefaultPhone = sitedetailDescPhone.Default.(string)
	// sitedetailDescAddress is the schema descriptor for address field.
	sitedetailDescAddress := sitedetailFields[6].Descriptor()
	// sitedetail.DefaultAddress holds the default value on creation for the address field.
	sitedetail.DefaultAddress = sitedetailDescAddress.Default.(string)
	// sitedetailDescFb is the schema descriptor for fb field.
	sitedetailDescFb := sitedetailFields[7].Descriptor()
	// sitedetail.DefaultFb holds the default value on creation for the fb field.
	sitedetail.DefaultFb = sitedetailDescFb.Default.(string)
	// sitedetailDescTw is the schema descriptor for tw field.
	sitedetailDescTw := sitedetailFields[8].Descriptor()
	// sitedetail.DefaultTw holds the default value on creation for the tw field.
	sitedetail.DefaultTw = sitedetailDescTw.Default.(string)
	// sitedetailDescWh is the schema descriptor for wh field.
	sitedetailDescWh := sitedetailFields[9].Descriptor()
	// sitedetail.DefaultWh holds the default value on creation for the wh field.
	sitedetail.DefaultWh = sitedetailDescWh.Default.(string)
	// sitedetailDescIg is the schema descriptor for ig field.
	sitedetailDescIg := sitedetailFields[10].Descriptor()
	// sitedetail.DefaultIg holds the default value on creation for the ig field.
	sitedetail.DefaultIg = sitedetailDescIg.Default.(string)
	// sitedetailDescID is the schema descriptor for id field.
	sitedetailDescID := sitedetailFields[0].Descriptor()
	// sitedetail.DefaultID holds the default value on creation for the id field.
	sitedetail.DefaultID = sitedetailDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[3].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[4].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[5].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[6].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[7].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescTermsAgreement is the schema descriptor for terms_agreement field.
	userDescTermsAgreement := userFields[8].Descriptor()
	// user.DefaultTermsAgreement holds the default value on creation for the terms_agreement field.
	user.DefaultTermsAgreement = userDescTermsAgreement.Default.(bool)
	// userDescIsEmailVerified is the schema descriptor for is_email_verified field.
	userDescIsEmailVerified := userFields[9].Descriptor()
	// user.DefaultIsEmailVerified holds the default value on creation for the is_email_verified field.
	user.DefaultIsEmailVerified = userDescIsEmailVerified.Default.(bool)
	// userDescIsStaff is the schema descriptor for is_staff field.
	userDescIsStaff := userFields[10].Descriptor()
	// user.DefaultIsStaff holds the default value on creation for the is_staff field.
	user.DefaultIsStaff = userDescIsStaff.Default.(bool)
	// userDescIsActive is the schema descriptor for is_active field.
	userDescIsActive := userFields[11].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the is_active field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
