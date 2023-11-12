package workspace

import (
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	user, err := middleware.ExtractUser(c)
	if err != nil {
		return err
	}

	workspaces, err := user.Access.ListWorkspaces()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(workspaces)
}