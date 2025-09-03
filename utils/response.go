package util

import "go-clean-code/constant"

type BodyResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func response(success bool, message string, data interface{}) BodyResponse {
	return BodyResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func ResponseSuccess(message string, data interface{}) BodyResponse {
	return response(true, message, data)
}

func ResponseNotFound(message string) BodyResponse {
	var newMessage string
	if message == "" {
		newMessage = constant.NotFound()
	} else {
		newMessage = message
	}
	return response(false, newMessage, nil)
}

func ResponseError(message string, data interface{}) BodyResponse {
	return response(false, message, data)
}
