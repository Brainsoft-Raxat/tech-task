package apperror

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type ErrorCode struct {
	code    int
	status  int
	message string
}

func NewErrorCode(code, status int, message string) ErrorCode {
	return ErrorCode{
		code:    code,
		status:  status,
		message: message,
	}
}

type ErrorInfo struct {
	Code             int    `json:"code"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	error
}

func NewErrorInfo(ctx context.Context, errorCode ErrorCode, developerMessage string) *ErrorInfo {
	return &ErrorInfo{
		Code:             errorCode.code,
		Status:           errorCode.status,
		Message:          errorCode.message,
		DeveloperMessage: developerMessage,
	}
}

func (e *ErrorInfo) Error() string {
	if e.error == nil {
		return fmt.Sprintf("%d %d %s %s", e.Code, e.Status, e.Message, e.DeveloperMessage)
	}

	return fmt.Sprintf("%d %d %s %s: %v", e.Code, e.Status, e.Message, e.DeveloperMessage, e.error)
}

func (e *ErrorInfo) Equal(err error) bool {
	appErr := AsErrorInfo(err)

	return appErr != nil && e.Code == appErr.Code
}

func (e *ErrorInfo) Cause() error {
	return e.error
}

func (e *ErrorInfo) Unwrap() error {
	return e.error
}

func (e *ErrorInfo) Wrap(err error) *ErrorInfo {
	if e.error != nil {
		err = errors.Wrap(err, e.error.Error())
	}

	appErr := &ErrorInfo{
		Code:             e.Code,
		Status:           e.Status,
		Message:          e.Message,
		DeveloperMessage: e.DeveloperMessage,
		error:            err,
	}

	return appErr
}

func (e *ErrorInfo) copy() *ErrorInfo {
	err := *e

	return &err
}

func (e *ErrorInfo) SetDeveloperMessage(message string) *ErrorInfo {
	err := e.copy()
	err.DeveloperMessage = message

	return err
}

func (e *ErrorInfo) SetMessage(message string) *ErrorInfo {
	err := e.copy()
	err.Message = message

	return err
}

func (e *ErrorInfo) SetHttpStatus(httpStatus int) *ErrorInfo {
	err := e.copy()
	err.Status = httpStatus

	return err
}

func AsErrorInfo(err error) *ErrorInfo {
	var target *ErrorInfo
	if errors.As(err, &target) {
		return target
	}

	return nil
}

func EqualWithErrorCode(err error, errorCode ErrorCode) bool {
	appErr := AsErrorInfo(err)

	return appErr != nil && appErr.Code == errorCode.code
}
