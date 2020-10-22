package utilities

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrWrongPassword    = errors.New("wrong password")
	ErrUsernameTaken    = errors.New("username was taken")
	ErrEmptyComment     = errors.New("empty comment")
	ErrTooShortPassword = errors.New("too short password")
	ErrTooShortUsername = errors.New("too short username")
	ErrEmptyPhone       = errors.New("empty phone")
	ErrEmptyEmail       = errors.New("empty email")
	ErrEmptyBirth       = errors.New("empty birthday")
)


func InternalServerError(writer http.ResponseWriter) {
	errorWrite(writer, http.StatusInternalServerError, "Internal server error")
}

func EmptyCommentError(writer http.ResponseWriter) {
	errorWrite(writer, http.StatusInternalServerError, "Empty comment")
}

func ParseFormError(writer http.ResponseWriter) {
	errorWrite(writer, http.StatusInternalServerError, "Parse form error")
}

func errorWrite(writer http.ResponseWriter, statusCode int, message string) {
	writer.WriteHeader(statusCode)
	_, _ = writer.Write([]byte(message))
}