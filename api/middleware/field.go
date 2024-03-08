package middleware

import (
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

func ExtractField(c *fiber.Ctx) (*types.User, *field.Field, error) {
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
