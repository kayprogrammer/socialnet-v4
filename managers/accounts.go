package managers

import (
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/otp"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

// ----------------------------------
// OTP MANAGEMENT
// --------------------------------
type UserManager struct {
}

func (obj UserManager) GetById(client *ent.Client, id uuid.UUID) *ent.User {
	u, _ := client.User.
		Query().
		Where(user.ID(id)).
		Only(Ctx)
	return u
}

func (obj UserManager) GetByRefreshToken(client *ent.Client, token string) *ent.User {
	u, _ := client.User.
		Query().
		Where(user.Refresh(token)).
		Only(Ctx)
	return u
}

func (obj UserManager) GetByEmail(client *ent.Client, email string) *ent.User {
	u, _ := client.User.
		Query().
		Where(user.Email(email)).
		Only(Ctx)
	return u
}

func (obj UserManager) GetByUsername(client *ent.Client, username string) *ent.User {
	u, _ := client.User.
		Query().
		Where(user.Username(username)).
		Only(Ctx)
	return u
}

func (obj UserManager) GetByUsernames(client *ent.Client, usernames []string, excludeOpts ...uuid.UUID) []*ent.User {
	usersQ := client.User.
		Query().
		Where(user.UsernameIn(usernames...))
	if len(excludeOpts) > 0 {
		usersQ = usersQ.Where(user.IDNEQ(excludeOpts[0]))
	}
	users := usersQ.AllX(Ctx)
	return users
}

func (obj UserManager) Create(client *ent.Client, userData schemas.RegisterUser, isStaff bool, isEmailVerified bool) *ent.User {
	username := UsernameGenerator(client, userData.FirstName, userData.LastName, nil)
	password := utils.HashPassword(userData.Password)
	u := client.User.Create().
		SetFirstName(userData.FirstName).
		SetLastName(userData.LastName).
		SetEmail(userData.Email).
		SetUsername(username).
		SetPassword(password).
		SetTermsAgreement(userData.TermsAgreement).
		SetIsStaff(isStaff).
		SetIsEmailVerified(isEmailVerified).
		SaveX(Ctx)
	return u
}

func (obj UserManager) GetOrCreate(client *ent.Client, userData schemas.RegisterUser, isEmailVerified bool, isStaff bool) *ent.User {
	user := obj.GetByEmail(client, userData.Email)
	if user == nil {
		// Create user
		user = obj.Create(client, userData, isStaff, isEmailVerified)
	}
	return user
}

func (obj UserManager) UpdateTokens(user *ent.User, access string, refresh string) *ent.User {
	u := user.Update().SetAccess(access).SetRefresh(refresh).SaveX(Ctx)
	return u
}

func (obj UserManager) DropData(client *ent.Client) {
	client.User.Delete().ExecX(Ctx)
}

// ----------------------------------
// OTP MANAGEMENT
// --------------------------------
type OtpManager struct {
}

func (obj OtpManager) GetOrCreate(client *ent.Client, userId uuid.UUID) *ent.Otp {
	code := utils.GetRandomInt(6)
	o := obj.GetByUserID(client, userId)

	// Create Otp
	if o == nil {
		o, _ = client.Otp.
			Create().
			SetUserID(userId).
			SetCode(code).
			Save(Ctx)
	} else {
		// Update the otp code
		o, _ = o.Update().SetCode(code).Save(Ctx)
	}
	return o
}

func (obj OtpManager) GetByUserID(client *ent.Client, userId uuid.UUID) *ent.Otp {
	o, _ := client.Otp.
		Query().
		Where(otp.UserID(userId)).
		Only(Ctx)
	return o
}

func (obj OtpManager) CheckExpiration(otpObj *ent.Otp) bool {
	currentTime := time.Now().UTC()
	diff := int64(currentTime.Sub(otpObj.UpdatedAt).Seconds())
	emailExpirySecondsTimeout := config.GetConfig().EmailOTPExpireSeconds
	return diff > emailExpirySecondsTimeout
}
