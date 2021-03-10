[![Go Report Card](https://goreportcard.com/badge/github.com/slazutkin/go-oauth-mailru)](https://goreportcard.com/report/github.com/slazutkin/go-oauth-mailru) [![GitHub license](https://img.shields.io/github/license/slazutkin/go-oauth-mailru)](https://github.com/slazutkin/go-oauth-mailru/blob/main/LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/slazutkin/go-oauth-mailru.svg)](https://pkg.go.dev/github.com/slazutkin/go-oauth-mailru)

# go-oauth-mailru
OAuth2 for Mail.ru

## Before start
[Mail.ru documentation on OAuth](https://help.mail.ru/developers/oauth)
## Usage
### Install:
```shell
go get github.com/slazutkin/go-oauth-mailru
```
### Import module:
```go
import oauth2mailru "github.com/slazutkin/go-oauth-mailru"
```
### Create instance of client:
```go
client := oauth2mailru.New(key, secret)
```
### Getting login URL
```go
url := client.LoginURL()
```

### Getting authorization token
```go
token, err := client.GetAuthorizationToken(authCode) (string, error)
```
### Getting user info
```go
user, err := client.GetUserInfo(token)
```
## Example
[See more details in complete example](https://github.com/slazutkin/go-oauth-mailru/tree/main/example)
