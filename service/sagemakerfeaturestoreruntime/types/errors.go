// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have permission to perform an action.
type AccessForbidden struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessForbidden) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessForbidden) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessForbidden) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AccessForbidden"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessForbidden) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal failure occurred. Try your request again. If the problem persists,
// contact Amazon Web Services customer support.
type InternalFailure struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailure) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalFailure"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A resource that is required to perform an action was not found.
type ResourceNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is currently unavailable.
type ServiceUnavailable struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailable) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ServiceUnavailable"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailable) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// There was an error validating your request.
type ValidationError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ValidationError"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
