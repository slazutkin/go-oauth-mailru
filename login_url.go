package oauth2mailru

import (
	"fmt"
	"net/url"
)

const loginURL = "https://oauth.mail.ru/login"

func (c *client) LoginURL(state string) string {

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", c.key)
	params.Add("client_secret", c.secret)
	params.Add("redirect_uri", c.redirectURL)

	if len(state) > 0 {
		params.Add("state", state)
	}

	p := params.Encode()

	return fmt.Sprintf("%s?%s", loginURL, p)
}
