package response

import "net/http"

const (
	Success               int = http.StatusOK
	Create                int = http.StatusCreated
	BadRequest            int = http.StatusBadRequest
	Unauthorized          int = http.StatusUnauthorized
	Forbidden             int = http.StatusForbidden
	NotFound              int = http.StatusNotFound
	InternalError         int = http.StatusInternalServerError
	TooManyRequests       int = http.StatusTooManyRequests
	EmailAlreadyExists    int = 20001
	UsernameAlreadyExists int = 20002
	WrongEmailOrPassword  int = 20003
	PasswordDoNotMatch    int = 20004
	VerifyEmailSuccess    int = 20005
)

var Msg = map[int]string{
	Success:               "Success",
	Create:                "Created successfully",
	BadRequest:            "Bad request",
	Unauthorized:          "Unauthorized access",
	Forbidden:             "Access forbidden",
	NotFound:              "Not found",
	InternalError:         "Internal server error",
	EmailAlreadyExists:    "Email already exists",
	UsernameAlreadyExists: "Username already exists",
	WrongEmailOrPassword:  "Wrong email or password",
	PasswordDoNotMatch:    "Password do not match",
	TooManyRequests:       "Too many requests",
	VerifyEmailSuccess:    "Verify email success",
}
