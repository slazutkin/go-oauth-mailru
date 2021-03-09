package oauth2mailru

import (
	"fmt"
	"testing"
)

func TestClientNewEmptyKey(t *testing.T) {
	key := ""
	secret := ""
	redirectURL := ""

	c, err := New(key, secret, redirectURL)

	if c != nil {
		t.Error("Expected nil pointer")
		return
	}

	if err == nil {
		t.Error("Expected non nil error")
		return
	}

	if err.Error() != fmt.Sprintf(missingParam, "key") {
		t.Errorf("Unexpected error: %v\n", err)
	}

}

func TestClientNewEmptySecret(t *testing.T) {
	key := "key"
	secret := ""
	redirectURL := ""

	c, err := New(key, secret, redirectURL)

	if c != nil {
		t.Error("Expected nil pointer")
		return
	}

	if err == nil {
		t.Error("Expected non nil error")
		return
	}

	if err.Error() != fmt.Sprintf(missingParam, "secret") {
		t.Errorf("Unexpected error: %v\n", err)
	}

}

func TestClientNewEmptyRedirectURL(t *testing.T) {
	key := "key"
	secret := "secret"
	redirectURL := ""

	c, err := New(key, secret, redirectURL)

	if c != nil {
		t.Error("Expected nil pointer")
		return
	}

	if err == nil {
		t.Error("Expected non nil error")
		return
	}

	if err.Error() != fmt.Sprintf(missingParam, "redirectURL") {
		t.Errorf("Unexpected error: %v\n", err)
	}

}
