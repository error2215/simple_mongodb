package models

import (
	"encoding/json"
)

type response struct {
	Success bool            `json:"success"`
	Error   *jsonError      `json:"error"`
	Data    json.RawMessage `json:"data"`
}

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(errorCode int, errorMsg string, data []byte) *response {
	if errorCode != 0 {
		return &response{
			Success: false,
			Error:   Error(errorCode, errorMsg),
			Data:    data,
		}
	}
	return &response{
		Success: true,
		Error:   nil,
		Data:    data,
	}
}
func (r *response) ToString() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

func Error(code int, message string) *jsonError {
	return &jsonError{
		Code:    code,
		Message: message,
	}
}
