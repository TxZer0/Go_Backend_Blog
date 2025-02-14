package response

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    Success,
		Message: Msg[Success],
		Data:    data,
	}
}

func NewCreateResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    Create,
		Message: Msg[Create],
		Data:    data,
	}
}

func NewVerifyEmailResponse() SuccessResponse {
	return SuccessResponse{
		Code:    Create,
		Message: Msg[Create],
		Data:    nil,
	}
}
