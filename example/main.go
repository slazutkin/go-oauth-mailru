package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	oauth2mailru "github.com/slazutkin/go-oauth-mailru"
)

func main() {
	key := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")
	redirectURL := "http://localhost:3000/oauth/mailru/callback"

	client, err := oauth2mailru.New(key, secret, redirectURL)

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		page := `
			<html>
				<head>
					<title>Mail.ru OAuth2 test</title>
				</head>
				<body>
					<a href="%s">Login via Mail.ru</a>
				</body>
			</html>
		`

		loginURL := client.LoginURL("some_state")

		fmt.Fprintf(w, page, loginURL)
	})

	mux.HandleFunc("/oauth/mailru/callback", func(w http.ResponseWriter, r *http.Request) {
		page := `
		<html>
		<head>
			<title>Mail.ru OAuth2 test</title>
		</head>
		<body>
			<h1>User info</h1>
			<div>
				<img src="%s">
				<ul>
					<li>Name: %s</li>
					<li>Email: %s</li>
				</ul>
			</div>
		</body>
		</html>
		`

		values := r.URL.Query()
		code := values.Get("code")

		token, err := client.GetAuthorizationToken(code)

		if err != nil {
			fmt.Fprintf(w, "Unexpected error: %s", err.Error())
			return
		}

		user, err := client.GetUserInfo(token)

		if err != nil {
			fmt.Fprintf(w, "Unexpected error: %s", err.Error())
			return
		}

		fmt.Fprintf(w, page,
			user.Image,
			fmt.Sprintf("%s %s", user.LastName, user.FirstName),
			user.Email,
		)
	})

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
