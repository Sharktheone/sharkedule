package user

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/user/access"
)

type User struct {
	User   types.User
	Access *access.Access
}

func from(user types.User) *User {
	return &User{
		User:   user,
		Access: &access.Access{Access: user.Access},
	}
}
