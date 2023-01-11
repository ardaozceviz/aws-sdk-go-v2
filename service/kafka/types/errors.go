// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Returns information about an error.
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returns information about an error.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returns information about an error.
type ForbiddenException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ForbiddenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returns information about an error.
type InternalServerErrorException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Returns information about an error.
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returns information about an error.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Returns information about an error.
type TooManyRequestsException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TooManyRequestsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returns information about an error.
type UnauthorizedException struct {
	Message *string

	ErrorCodeOverride *string

	InvalidParameter *string

	noSmithyDocumentSerde
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UnauthorizedException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
