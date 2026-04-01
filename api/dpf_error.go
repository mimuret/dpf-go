package api

import (
	"fmt"
)

// DPFError is an interface for common DPF API errors.
type DPFError interface {
	error
	GetRequestId() string
	GetErrorType() string
	GetErrorMessage() string
}

// Ensure all error models implement DPFError
var _ DPFError = &BadRequest{}
var _ DPFError = &TargetBadRequest{}
var _ DPFError = &NotFound{}
var _ DPFError = &TooManyRequests{}
var _ DPFError = &SystemError{}
var _ DPFError = &ServiceUnavailable{}
var _ DPFError = &GatewayTimeout{}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *TargetBadRequest) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *TooManyRequests) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *ServiceUnavailable) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}

func (e *GatewayTimeout) Error() string {
	return fmt.Sprintf("DPF Error: %s (RequestID: %s)", e.ErrorMessage, e.RequestId)
}
