package internal

import (
	"os"
	"strings"
)

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
