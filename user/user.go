package user

import (
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/user/access"
)

type User struct {
	User   *ktypes.User
	Access *access.Access
}

func From(user *ktypes.User) *User {
	return &User{
		User:   user,
		Access: &access.Access{Access: user.Access},
	}
}
