package rsgohttpresponse

import (
	json "encoding/json"
	http "net/http"
	reflect "reflect"
)

var (
	defaultHeaders = map[string]string{
		"Content-Type": "application/json",
	}

	mapping = map[string]CustomHTTPResponse{

		CodeSuccess: CustomHTTPResponse{
			StatusCode: http.StatusOK,
			Headers:    defaultHeaders,
		},

		CodeBadLogin: CustomHTTPResponse{
			StatusCode: http.StatusUnauthorized,
			Headers:    defaultHeaders,
		},

		CodeInvalidToken: CustomHTTPResponse{
			StatusCode: http.StatusUnauthorized,
			Headers:    defaultHeaders,
		},
		CodeAlreadyExists: CustomHTTPResponse{
			StatusCode: http.StatusConflict,
			Headers:    defaultHeaders,
		},
		CodeDoesNotExist: CustomHTTPResponse{
			StatusCode: http.StatusNotFound,
			Headers:    defaultHeaders,
		},
		CodeInvalidJSON: CustomHTTPResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    defaultHeaders,
		},
		CodeUpdated: CustomHTTPResponse{
			StatusCode: http.StatusOK,
			Headers:    defaultHeaders,
		},
		CodeValidationFailed: CustomHTTPResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    defaultHeaders,
		},
		CodeInternalError: CustomHTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    defaultHeaders,
		},
		CodeNotImplemented: CustomHTTPResponse{
			StatusCode: http.StatusNotImplemented,
			Headers:    defaultHeaders,
		},
	}
)

// CustomHTTPResponse : Custom HTTP Response sent back to client
type CustomHTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       CustomHTTPResponseBody
}

// CustomHTTPResponseBody : Custom HTTP Response body to send to client
type CustomHTTPResponseBody struct {
	ResponseDetails ResponseDetails `json:"details"`
	Content         interface{}     `json:"content"`
}

// WriteResponse : Return an HTTP response with chosen status
func WriteResponse(content interface{}, responseDetails *ResponseDetails, w http.ResponseWriter) {

	var responseBody CustomHTTPResponseBody

	responseBody.ResponseDetails = *responseDetails
	responseBody.Content = content

	// Retrieve response settings
	settings := mapping[responseDetails.Code]

	// Write Headers
	for _, key := range reflect.ValueOf(settings.Headers).MapKeys() {

		w.Header().Set(key.String(), settings.Headers[key.String()])
	}

	// Write Status Code
	w.WriteHeader(settings.StatusCode)
	res, _ := json.Marshal(responseBody)
	w.Write(res)
}
