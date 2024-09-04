package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ApiResponse(w http.ResponseWriter, status int, body interface{}) *http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	stringBody, err := json.Marshal(body)
	if err != nil {
		fmt.Print("Error While Marshal")
	}

	w.Write(stringBody)

	return &w
}

func GetLastString(query string, delimiter string) string {
	parts := strings.Split(query, delimiter)
	lastPart := parts[len(parts)-1]
	return lastPart
}
