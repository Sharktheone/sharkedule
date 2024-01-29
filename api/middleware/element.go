package middleware

import (
	"github.com/Sharktheone/sharkedule/element"
	"github.com/gofiber/fiber/v2"
)

func ExtractElement(c *fiber.Ctx) (error, *element.Element) {
	elementUUID := c.Params("element")

	elem, err := element.GetElement(elementUUID) //TODO
	if err != nil {
		return err, nil
	}

	return nil, elem
}
