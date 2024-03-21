package element

import (
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func FieldList(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	return c.JSON(e.ListFields())
}

func FieldInfo(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	field := c.Params("field")

	return c.JSON(e.GetField(field))
}

func FieldUpdate(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	field := c.Params("field")

	payload := new(struct {
		Value string `json:"value"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return e.UpdateField(field, payload.Value)
}

func FieldDeleteScoped(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	field := c.Params("field")

	return e.DeleteScopedField(field)
}

func FieldCreateScoped(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Name string `json:"name"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return e.CreateScopedField(payload.Name)
}
