package errors

import "net/http"

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

var (
	FailedGeneratingToken = Error{Message: "Failed to generate token", Status: http.StatusInternalServerError}
	FailedHashingPassword = Error{Message: "Failed to hash password", Status: http.StatusInternalServerError}
	InvalidPassword       = Error{Message: "Invalid password", Status: http.StatusUnauthorized}
	InvalidRequestBody    = Error{Message: "Invalid request body", Status: http.StatusBadRequest}
	MethodNotAllowed      = Error{Message: "Method not allowed", Status: http.StatusMethodNotAllowed}
	UserAlreadyExists     = Error{Message: "User already exists", Status: http.StatusForbidden}
	UserNotFound          = Error{Message: "User not found", Status: http.StatusNotFound}
	InvalidClaims         = Error{Message: "Invalid claims", Status: http.StatusUnauthorized}
	FailedParsingToken    = Error{Message: "Failed to parse token", Status: http.StatusInternalServerError}
)
