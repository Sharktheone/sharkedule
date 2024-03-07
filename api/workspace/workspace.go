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

func ListWithFields(c *fiber.Ctx) error {
	return nil
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

func Delete(c *fiber.Ctx) error {
	user, err := middleware.ExtractUser(c)
	if err != nil {
		return err
	}

	uuid := c.Params("uuid")

	err = user.Access.DeleteWorkspace(uuid)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
