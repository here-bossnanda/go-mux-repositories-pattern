package response

import (
	"api/constants"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type IResponse interface {
	Success(w http.ResponseWriter)
	Error(w http.ResponseWriter, message ...string)
}

// Success is response with json format
func (res Response) Success(w http.ResponseWriter) {

	res.Status = http.StatusOK
	res.Message = "OK"

	json.NewEncoder(w).Encode(res)
}

// Error is response with json format
func (res Response) Error(w http.ResponseWriter, message ...string) {
	s, m := generateError(message)
	if res.Status == 0 {
		res.Status = s
	}
	res.Message = m
	json.NewEncoder(w).Encode(res)
}

// generateError is used to map error message
func generateError(messages []string) (int, string) {
	var status int

	if messages == nil {
		messages = append(messages, constants.ThereSomethingWrong)
	}

	switch messages[0] {
	case constants.NoResult:
		messages = []string{}
		messages = append(messages, constants.NoRecordFound)
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}

	message := fmt.Sprintf(strings.Join(messages, "\n"))

	return status, message
}
