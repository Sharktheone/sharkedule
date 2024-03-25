package middleware

import (
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

func ExtractElement(c *fiber.Ctx) (types.User, types.Element, error) {
	workspace := c.Params("workspace")
	elementUUID := c.Params("element")

	u, err := ExtractUser(c)
	if err != nil {
		return u, nil, err
	}

	elem, err := u.GetAccess().GetElement(workspace, elementUUID)
	if err != nil {
		return u, nil, err
	}

	elem.SetUser(u.GetUUID())

	return u, elem, nil
}

func ExtractElementAccess(c *fiber.Ctx) (types.Access, error) {
	workspace := c.Params("workspace")
	elementUUID := c.Params("element")

	u, err := ExtractUser(c)
	if err != nil {
		return nil, err
	}

	access := u.GetAccess()

	access.SetWorkspace(workspace)
	access.SetElement(elementUUID)

	return access, nil
}
