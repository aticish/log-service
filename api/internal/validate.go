package internal

import (
	"encoding/json"
	"net"
	"os"
	"strings"
)

// CheckToken controls a bearer token is valid or not
func CheckToken(token string) error {
	token = strings.TrimSpace(token)
	tokenParts := strings.SplitN(token, " ", 2)

	// Check if token has the expected format "Bearer TOKEN"
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ErrorInvalidToken
	}

	// Extract the actual token value
	tokenValue := strings.TrimSpace(tokenParts[1])

	// Compare with environment variable
	if strings.Compare(os.Getenv("LOGAPI_TOKEN"), tokenValue) == 0 {
		return nil
	}
	return ErrorInvalidToken
}

// CheckIP controls an IP address is valid or not
func CheckIP(ip string) bool {
	validated := net.ParseIP(ip)
	if validated == nil {
		return false
	}
	return true
}

// CheckSeverity controls a severity is valid or not
func CheckSeverity(severity string) bool {
	severity = strings.ToLower(severity)
	switch severity {
	case "emergency", "alert", "critical", "error", "warning", "notice", "info", "debug":
		return true
	}
	return false
}

// CheckContent controls a content property value is valid or not
func CheckContent(content map[string]any) bool {
	byteContent, err := json.Marshal(content)
	if err != nil {
		return false
	}
	if !json.Valid(byteContent) {
		return false
	}
	return true
}
