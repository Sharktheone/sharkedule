package workspace

import (
	"github.com/Sharktheone/sharkedule/api/middleware"
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

func Info(c *fiber.Ctx) error {
	user, err := middleware.ExtractUser(c)
	if err != nil {
		return err
	}

	workspace, err := user.Access.WorkspaceInfo()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(workspace)
}
