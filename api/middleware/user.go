package middleware

import (
	"github.com/Sharktheone/sharkedule/user"
	"github.com/gofiber/fiber/v2"
)

func ExtractUser(c *fiber.Ctx) (*user.User, error) {
	//TODO: maybe use jwt or something else token based
	return nil, nil
}
