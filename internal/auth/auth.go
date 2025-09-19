package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey, it extracts an api key from the headers of an http request
// Example, Authorization: APIKey {insert APIKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("maloformed auth header")
	}

	if vals[0] != "APIKey" {
		return "", errors.New("maloformed first part of auth header")
	}

	return vals[1], nil
}

func GetName(headers http.Header) (string, error) {
	val := headers.Get("Name")
	if val == "" {
		return "", errors.New("no name info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("maloformed name header")
	}

	if vals[0] != "Name" {
		return "", errors.New("maloformed first part of name header")
	}

	return vals[1], nil
}
