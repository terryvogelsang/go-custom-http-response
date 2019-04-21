package rsgohttpresponse

import (
	uuid "github.com/satori/go.uuid"
)

const (
	CodeSuccess           = "SUCCESS"
	CodeUpdated           = "UPDATED"
	CodeBadLogin          = "BAD-LOGIN"
	CodeInvalidToken      = "INVALID-TOKEN"
	CodeAlreadyExists     = "ALREADY-EXISTS"
	CodeDoesNotExist      = "DOES-NOT-EXIST"
	CodeInvalidJSON       = "INVALID-JSON"
	CodeValidationFailed  = "VALIDATION-FAILED"
	CodeInternalError     = "INTERNAL-ERROR"
	CodeNotImplemented    = "NOT-IMPLEMENTED"
)

var (
	messageCodeMapping = map[string]string{
		CodeBadLogin:          "Bad username or password",
		CodeInvalidToken:      "Invalid Token",
		CodeSuccess:           "Success",
		CodeAlreadyExists:     "Resource already exists",
		CodeInvalidJSON:       "Invalid JSON Format",
		CodeUpdated:           "Resource updated",
		CodeValidationFailed:  "Field validation failed",
		CodeInternalError:     "An internal error occured",
		CodeDoesNotExist:      "Resource does not exist",
		CodeNotImplemented:    "Feature is not implemented yet",
	}
)

// ResponseDetails informations about log entry
type ResponseDetails struct {
	LogID   string   `json:"logID"`
	Session string   `json:"session,omitempty"`
	Service string   `json:"service"`
	Action  string   `json:"action"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Fields  []string `json:"failedFields,omitempty"`
	Debug   string   `json:"debug,omitempty"`
}

// NewResponseDetailsWithSession return a ResponseDetails struct instance with a session
func NewResponseDetailsWithSession(session string, service string, action string, code string) *ResponseDetails {

	uuid := uuid.NewV4()

	return &ResponseDetails{
		LogID:   uuid.String(),
		Session: session,
		Service: service,
		Action:  action,
		Code:    code,
		Message: messageCodeMapping[code],
	}
}

// NewResponseDetails return a ResponseDetails struct instance
func NewResponseDetails(service string, action string, code string) *ResponseDetails {

	uuid := uuid.NewV4()

	return &ResponseDetails{
		LogID:   uuid.String(),
		Service: service,
		Action:  action,
		Code:    code,
		Message: messageCodeMapping[code],
	}
}

// NewResponseDetailsWithDebug return a ResponseDetails struct instance with full error for debugging purpose
func NewResponseDetailsWithDebug(debug string, service string, action string, code string) *ResponseDetails {

	uuid := uuid.NewV4()

	return &ResponseDetails{
		LogID:   uuid.String(),
		Service: service,
		Action:  action,
		Code:    code,
		Message: messageCodeMapping[code],
		Debug:   debug,
	}
}

// NewResponseDetailsWithFields return a ResponseDetails struct instance with failed fields
func NewResponseDetailsWithFields(fields []string, service string, action string, code string) *ResponseDetails {

	uuid := uuid.NewV4()

	return &ResponseDetails{
		LogID:   uuid.String(),
		Service: service,
		Action:  action,
		Code:    code,
		Message: messageCodeMapping[code],
		Fields:  fields,
	}
}
