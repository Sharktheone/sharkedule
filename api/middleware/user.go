package middleware

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/user"
)

func ExtractUser(c *fiber.Ctx) (*user.User, error) {
	//TODO: maybe use jwt or something else token based
	u, err := db.DB.GetUser("69b78d55-2058-4440-a7e7-183383f3c0dd") //TODO: just for testing
	if err != nil {
		return nil, err
	}

	return user.From(u), nil
}
