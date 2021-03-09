package oauth2mailru

import "fmt"

type client struct {
	key         string
	secret      string
	redirectURL string
}

const (
	missingParam = "Parameter %s cannot be empty"
)

// New ...
func New(key, secret, redirectURL string) (MailruClient, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf(missingParam, "key")
	}

	if len(secret) == 0 {
		return nil, fmt.Errorf(missingParam, "secret")
	}

	if len(redirectURL) == 0 {
		return nil, fmt.Errorf(missingParam, "redirectURL")
	}

	return &client{key, secret, redirectURL}, nil
}
