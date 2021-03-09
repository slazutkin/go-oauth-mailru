package oauth2mailru

import "context"

// MailruClient ...
type MailruClient interface {
	LoginURL(state string) string
	GetAuthorizationToken(authCode string) (string, error)
	GetUserInfo(ctx context.Context, token string) (*UserInfo, error)
}
