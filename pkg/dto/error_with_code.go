package dto

import (
	"fmt"
)

type ErrorWithCode struct {
	Code    int
	Message string
}

const (
	// ERROR code
	ErrorCodeOK              = 0x0
	ErrorCodeFoundChatServer = 0x00010001
	ErrorCodeNotFundCmd      = 0x00020001
	ErrorCodeInvalidParam    = 0x00020002
	ErrorCodeInvalidStatus   = 0x00020003
	ErrorCodeClientNotFound  = 0x00020004
	ErrorCodeServerError     = 0x00020005

	// request Id
	RequestIdFromServer = "0"
)

var (
	ErrorFoundChatServer = ErrorWithCode{ErrorCodeFoundChatServer, "cannot found chat server"}
	ErrorNotFoundCmd     = ErrorWithCode{ErrorCodeNotFundCmd, "cmd cannot found"}
	ErrorOK              = ErrorWithCode{ErrorCodeOK, "OK"}
	ErrorInvalidParam    = ErrorWithCode{ErrorCodeInvalidParam, "invalid param"}
	ErrorClientNotFound  = ErrorWithCode{ErrorCodeClientNotFound, "client not found"}
	ErrorServerError     = ErrorWithCode{ErrorCodeServerError, "server error"}
)

func (err ErrorWithCode) Error() string {
	return fmt.Sprintf("error code: %d, msg: %s", err.Code, err.Message)
}
