package auth

import (
	"net/http"
	"testing"
)

const authorization = "Authorization"

func TestGetApiKeyNoAuthorizationHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetApiKeyEmptyAuthorizationHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		"Authorization": {""},
	})

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetAPIKeyBadHeaderValue(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		"Authorization": {"bad-value"},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetApiKeyShouldReturnKey(t *testing.T) {
	expected := "fake-api-key"
	got, err := GetAPIKey(http.Header{
		"Authorization": {"ApiKey " + expected},
	})
	if err != nil {
		t.Fatal("unexpected error")
	}

	if got != expected {
		t.Fatalf("values mismatch expected: %s got: %s", expected, got)
	}
}
