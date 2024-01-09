package utils

type ErrorResponse struct {
	Status				string				`json:"status"`
	Code				string				`json:"code"`
	Message				string				`json:"message"`
	Data				*map[string]string	`json:"data,omitempty"`
}

func (obj ErrorResponse) Init() ErrorResponse {
	if obj.Status == "" {
		obj.Status = "failure"
	}
	return obj
}

type ErrorCodeStruct struct {
	UNAUTHORIZED_USER		string
	NETWORK_FAILURE			string
	SERVER_ERROR			string
	INVALID_ENTRY			string
	INCORRECT_EMAIL			string
	INCORRECT_OTP			string
	EXPIRED_OTP				string
	INVALID_AUTH			string
	INVALID_TOKEN			string
	INVALID_CREDENTIALS		string
	UNVERIFIED_USER			string
	NON_EXISTENT			string
	INVALID_OWNER			string
	INVALID_PAGE			string
	INVALID_VALUE			string
	NOT_ALLOWED				string
	INVALID_DATA_TYPE		string
}

var ErrorCode = ErrorCodeStruct{
	UNAUTHORIZED_USER:	"unauthorized_user",
	NETWORK_FAILURE:	"network_failure",
	SERVER_ERROR:	"server_error",
	INVALID_ENTRY:	"invalid_entry",
	INCORRECT_EMAIL:	"incorrect_email",
	INCORRECT_OTP:	"incorrect_otp",
	EXPIRED_OTP:	"expired_otp",
	INVALID_AUTH:	"invalid_auth",
	INVALID_TOKEN:	"invalid_token",
	INVALID_CREDENTIALS:	"invalid_credentials",
	UNVERIFIED_USER:	"unverified_user",
	NON_EXISTENT:	"non_existent",
	INVALID_OWNER:	"invalid_owner",
	INVALID_PAGE:	"invalid_page",
	INVALID_VALUE:	"invalid_value",
	NOT_ALLOWED:	"not_allowed",
	INVALID_DATA_TYPE:	"invalid_data_type",
}