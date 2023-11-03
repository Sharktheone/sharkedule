package user

import (
	"github.com/Sharktheone/sharkedule/user/access"
	"github.com/Sharktheone/sharkedule/user/email"
	"github.com/Sharktheone/sharkedule/user/mfa"
	"github.com/Sharktheone/sharkedule/user/oauth"
	"github.com/Sharktheone/sharkedule/user/settings"
)

type User struct {
	UUID     string
	Username string
	Password string
	OAuth    oauth.OAuth
	MFA      mfa.MFA
	Email    email.EMail
	Access   access.Access
	Settings settings.Settings
}
