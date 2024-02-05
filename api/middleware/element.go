package middleware

import (
	"github.com/Sharktheone/sharkedule/element"
	"github.com/Sharktheone/sharkedule/user"
	"github.com/gofiber/fiber/v2"
)

func ExtractElement(c *fiber.Ctx) (*user.User, *element.Element, error) {
	workspace := c.Params("workspace")
	elementUUID := c.Params("element")

	u, err := ExtractUser(c)
	if err != nil {
		return u, nil, err
	}

	elem, err := u.Access.GetElement(workspace, elementUUID)
	if err != nil {
		return u, nil, err
	}

	elem.SetUser(u)

	return u, elem, nil
}
