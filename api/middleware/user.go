package middleware

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

func ExtractUser(c *fiber.Ctx) (types.User, error) {
	//TODO: maybe use jwt or something else token based
	return db.DB.GetUser("69b78d55-2058-4440-a7e7-183383f3c0dd") //TODO: just for testing
}
