package oauth2mailru

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// UserInfo ...
type UserInfo struct {
	ID        string `json:"id"`
	ClientID  string `json:"client_id"`
	Gender    string `json:"gender"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Locale    string `json:"locale"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
}

const userInfoURL = "https://oauth.mail.ru/userinfo"

func (c *client) GetUserInfo(ctx context.Context, token string) (*UserInfo, error) {
	params := url.Values{}
	params.Add("access_token", token)

	r, err := http.NewRequest(http.MethodPost, userInfoURL, strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}

	r.SetBasicAuth(c.key, c.secret)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))

	client := http.Client{}

	res, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var u UserInfo

	if err := json.Unmarshal(response, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
