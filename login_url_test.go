package oauth2mailru

import (
	"fmt"
	"net/url"
	"testing"
)

func TestLoginUrlWithState(t *testing.T) {
	key := "testkey"
	secret := "testsecret"
	redirectURL := "http://localhost/callback"
	state := "some_state"

	params := url.Values{}

	params.Set("response_type", "code")
	params.Set("client_id", key)
	params.Set("client_secret", secret)
	params.Set("redirect_uri", redirectURL)
	params.Set("state", state)

	expected := fmt.Sprintf("%s?%s", loginURL, params.Encode())

	c, err := New(key, secret, redirectURL)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	got := c.LoginURL(state)

	if expected != got {
		t.Errorf("\nExpected %s\nGot %s\n", expected, got)
	}
}
