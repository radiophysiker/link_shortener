package utils

import (
	"net/http"
)

type Logger interface {
	Error(format string, args ...any)
}

// WriteErrorWithCannotWriteResponse Write error message to the log and sets the status code to 500,
// if w.Write return err.
func WriteErrorWithCannotWriteResponse(w http.ResponseWriter, err error, log Logger) {
	log.Error("cannot write response: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
}
