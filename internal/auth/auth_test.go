package auth

import (
	"net/http"
	"testing"
)

const authorizationHeader = "Authorization"

func TestGetAPIKeyNoAuthorizationHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestGetAPIKeyEmptyAuthorizationHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		authorizationHeader: {""},
	})

	if err == nil {
		t.Fatal("error expected")
	}
}

func TestGetAPIKeyBadHeaderValue(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		authorizationHeader: {"bad-value"},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetAPIKeyShouldReturnKey(t *testing.T) {
	expected := "fake-api-key"
	got, err := GetAPIKey(http.Header{
		authorizationHeader: {"ApiKey " + expected},
	})
	if err != nil {
		t.Fatal("unexpected error")
	}

	if got != expected {
		t.Fatalf("values mismatch expected: %s got: %s", expected, got)
	}
}
