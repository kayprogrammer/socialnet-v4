package schemas

// REQUEST BODY SCHEMAS
type RegisterUser struct {
	FirstName      string `json:"first_name" validate:"required,max=50" example:"John"`
	LastName       string `json:"last_name" validate:"required,max=50" example:"Doe"`
	Email          string `json:"email" validate:"required,min=5,email" example:"johndoe@email.com"`
	Password       string `json:"password" validate:"required,min=8,max=50" example:"strongpassword"`
	TermsAgreement bool   `json:"terms_agreement" validate:"eq=true"`
}

type EmailRequestSchema struct {
	Email string `json:"email" validate:"required,min=5,email" example:"johndoe@email.com"`
}

// RESPONSE BODY SCHEMAS
type RegisterResponseSchema struct {
	ResponseSchema
	Data EmailRequestSchema `json:"data"`
}
