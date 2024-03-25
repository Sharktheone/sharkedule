package middleware

import (
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

func ExtractWorkspace(c *fiber.Ctx) (types.User, types.Workspace, error) {
	workspace := c.Params("workspace")

	u, err := ExtractUser(c)
	if err != nil {
		return u, nil, err
	}

	w, err := u.GetAccess().GetWorkspace(workspace)
	if err != nil {
		return u, nil, err
	}

	return u, w, nil
}

func ExtractWorkspaceAccess(c *fiber.Ctx) (types.Access, error) {
	workspace := c.Params("workspace")

	u, err := ExtractUser(c)
	if err != nil {
		return nil, err
	}

	access := u.GetAccess()

	access.SetWorkspace(workspace)

	return access, nil
}
