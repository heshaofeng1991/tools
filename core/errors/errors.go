package errors

import "core/response"

func New(code int, msg string) error {
	return &response.Codes{
		Code: code,
		Msg:  msg,
	}
}
