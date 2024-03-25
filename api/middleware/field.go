package middleware

import (
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

func ExtractField(c *fiber.Ctx) (types.User, types.Field, error) {
	workspace := c.Params("workspace")
	fieldUUID := c.Params("field")
	u, err := ExtractUser(c)
	if err != nil {
		return nil, nil, err
	}

	f, err := u.GetAccess().GetField(workspace, fieldUUID)
	if err != nil {
		return nil, nil, err
	}

	return u, f, nil
}

func ExtractFieldAccess(c *fiber.Ctx) (types.Access, error) {
	workspace := c.Params("workspace")
	fieldUUID := c.Params("field")
	u, err := ExtractUser(c)
	if err != nil {
		return nil, err
	}

	access := u.GetAccess()

	access.SetWorkspace(workspace)
	access.SetField(fieldUUID)

	return access, nil
}

func ExtractElementField(c *fiber.Ctx) (types.User, types.Element, types.Field, error) {
	workspace := c.Params("workspace")
	elementUUID := c.Params("element")
	fieldUUID := c.Params("field")
	u, err := ExtractUser(c)
	if err != nil {
		return nil, nil, nil, err
	}

	e, err := u.GetAccess().GetElement(workspace, elementUUID)
	if err != nil {
		return nil, nil, nil, err
	}

	f, err := u.GetAccess().GetField(workspace, fieldUUID)
	if err != nil {
		return nil, nil, nil, err
	}

	return u, e, f, nil
}

func ExtractElementFieldAccess(c *fiber.Ctx) (types.Access, error) {
	workspace := c.Params("workspace")
	elementUUID := c.Params("element")
	fieldUUID := c.Params("field")
	u, err := ExtractUser(c)
	if err != nil {
		return nil, err
	}

	access := u.GetAccess()

	access.SetWorkspace(workspace)
	access.SetElement(elementUUID)
	access.SetField(fieldUUID)

	return access, nil
}
