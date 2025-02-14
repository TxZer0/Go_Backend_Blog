package response

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBadRequest() ErrorResponse {
	return ErrorResponse{
		Code:    BadRequest,
		Message: Msg[BadRequest],
	}
}

func NewUnauthorized() ErrorResponse {
	return ErrorResponse{
		Code:    Unauthorized,
		Message: Msg[Unauthorized],
	}
}

func NewForbidden() ErrorResponse {
	return ErrorResponse{
		Code:    Forbidden,
		Message: Msg[Forbidden],
	}
}

func NewNotFound() ErrorResponse {
	return ErrorResponse{
		Code:    NotFound,
		Message: Msg[NotFound],
	}
}

func NewInternalError() ErrorResponse {
	return ErrorResponse{
		Code:    InternalError,
		Message: Msg[InternalError],
	}
}

func NewWrongEmailOrPassword() ErrorResponse {
	return ErrorResponse{
		Code:    WrongEmailOrPassword,
		Message: Msg[WrongEmailOrPassword],
	}
}

func NewEmailAlreadyExists() ErrorResponse {
	return ErrorResponse{
		Code:    EmailAlreadyExists,
		Message: Msg[EmailAlreadyExists],
	}
}

func NewUsernameAlreadyExists() ErrorResponse {
	return ErrorResponse{
		Code:    UsernameAlreadyExists,
		Message: Msg[UsernameAlreadyExists],
	}
}

func NewPasswordDoNotMatch() ErrorResponse {
	return ErrorResponse{
		Code:    PasswordDoNotMatch,
		Message: Msg[PasswordDoNotMatch],
	}
}

func NewTooManyRequests() ErrorResponse {
	return ErrorResponse{
		Code:    TooManyRequests,
		Message: Msg[TooManyRequests],
	}
}
