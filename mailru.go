package oauth2mailru

// MailruClient ...
type MailruClient interface {
	LoginURL(state string) string
	GetAuthorizationToken(authCode string) (string, error)
	GetUserInfo(token string) (*UserInfo, error)
}
