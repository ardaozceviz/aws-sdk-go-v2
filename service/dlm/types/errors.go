// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The service failed in an unexpected way.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Bad request. The request is missing required parameters or has invalid
// parameters.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Code                        *string
	RequiredParameters          []string
	MutuallyExclusiveParameters []string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because a limit was exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code         *string
	ResourceType *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A requested resource was not found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code         *string
	ResourceType *string
	ResourceIds  []string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
