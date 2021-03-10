package oauth2mailru

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const authTokenURL = "https://oauth.mail.ru/token"

type authTokenResponse struct {
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *client) GetAuthorizationToken(authCode string) (string, error) {
	params := url.Values{}
	params.Add("code", authCode)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", c.redirectURL)

	r, err := http.NewRequest(http.MethodPost, authTokenURL, strings.NewReader(params.Encode()))

	if err != nil {
		return "", err
	}

	r.SetBasicAuth(c.key, c.secret)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))

	client := http.Client{}

	res, err := client.Do(r)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var auth authTokenResponse

	if err := json.Unmarshal(response, &auth); err != nil {
		return "", err
	}

	return auth.AccessToken, nil
}
