package managers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/otp"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type UserManager struct {
}

func (obj UserManager) GetById(client *ent.Client, id uuid.UUID) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(Ctx)
	if err != nil {
		fmt.Printf("failed querying user by id: %v\n", err)
		return nil, nil
	}
	return u, nil
}

func (obj UserManager) GetByEmail(client *ent.Client, email string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Email(email)).
		Only(Ctx)
	if err != nil {
		fmt.Printf("failed querying user by email: %v\n", err)
		return nil, nil
	}
	return u, nil
}

func (obj UserManager) Create(client *ent.Client, userData schemas.RegisterUser) (*ent.User, error) {
	username := UsernameGenerator(client, userData.FirstName, userData.LastName, nil, nil)
	password := utils.HashPassword(userData.Password)

	u, err := client.User.
		Create().
		SetFirstName(userData.FirstName). 
		SetLastName(userData.LastName).
		SetEmail(userData.Email). 
		SetUsername(username). 
		SetPassword(password).
		SetTermsAgreement(userData.TermsAgreement). 
		Save(Ctx)
	if err != nil {
		fmt.Printf("failed creating user: %v\n", err)
		return nil, nil
	}
	return u, nil
}

type OtpManager struct {
}

func (obj OtpManager) GetOrCreate(client *ent.Client, userId uuid.UUID) *ent.Otp {
	code := utils.GetRandomInt(6)
	o, _ := obj.GetByUserID(client, userId)
	
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

func (obj OtpManager) GetByUserID(client *ent.Client, userId uuid.UUID) (*ent.Otp, error) {
	o, err := client.Otp.
		Query().
		Where(otp.UserID(userId)).
		Only(Ctx)
	if err != nil {
		fmt.Printf("failed querying otp by user id: %v\n", err)
		return nil, nil
	}
	return o, nil
}

func (obj OtpManager) CheckExpiration(otpObj *ent.Otp) bool {
	currentTime := time.Now().UTC()
	diff := int64(currentTime.Sub(otpObj.UpdatedAt).Seconds())
	emailExpirySecondsTimeout := config.GetConfig().EmailOTPExpireSeconds
	return diff > emailExpirySecondsTimeout
}