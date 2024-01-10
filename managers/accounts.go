package managers

import (
	"fmt"

	"github.com/google/uuid"
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
	u, err := client.User.
		Create().
		SetFirstName(userData.FirstName). 
		SetLastName(userData.LastName).
		SetEmail(userData.Email). 
		SetUsername(username). 
		SetPassword(userData.Password).
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
	o, _ := client.Otp.
		Query().
		Where(otp.UserID(userId)).
		Only(Ctx)
	
	// Create Otp
	if o == nil {
		o, _ = client.Otp.
			Create().
			SetUserID(userId). 
			SetCode(code).
			Save(Ctx)
	} else {
		// Update the otp code
		o.Update().SetCode(code).Save(Ctx)
	}
	return o
}