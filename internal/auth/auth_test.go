package auth

import (
	"net/http"
	"testing"
)

func TestNoAuthorization(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "")
	result, err := GetAPIKey(header)

	if err != ErrNoAuthHeaderIncluded && result != "" {
		t.Fatalf("expected: \"\", got: %v", result)
	}
}

func TestNoHeader(t *testing.T) {
	header := make(http.Header)
	result, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded && result != "" {
		t.Fatalf("expected: \"\", got: %v", result)
	}
}

func TestCorrectApi(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "ApiKey mytestkey")
	result, err := GetAPIKey(header)
	if result != "mytestkey" || err != nil {
		t.Fatalf("expected: \"\", got: %v", result)
	}
}
