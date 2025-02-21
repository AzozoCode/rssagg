package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API Key from the headers
// of an HTTP request
// Example: Authorization : ApiKey {insert api key here}
// Returns the API Key or an error
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {

		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {

		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {

		return "", errors.New("unsupported auth type")
	}

	return vals[1], nil
}
