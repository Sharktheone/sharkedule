package field

import (
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	return c.JSON(w.ListFields())
}

func Info(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	field := c.Params("field")

	info, err := w.InfoField(field)
	if err != nil {
		return err

	}

	return c.JSON(info)
}

func Create(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Name string `json:"name"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	uuid, err := w.CreateField(payload.Name)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"uuid": uuid})
}

func Delete(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	field := c.Params("field")

	uuid, err := w.DeleteField(field)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"uuid": uuid})
}

func ListLinked(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	return c.JSON(w.ListLinkedFields())
}

func Link(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Field string `json:"field"`
		From  string `json:"from"`
		To    string `json:"to"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	uuid, err := w.LinkField(payload.Field, payload.From, payload.To)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"uuid": uuid})
}

func Unlink(c *fiber.Ctx) error {
	_, w, err := middleware.ExtractWorkspace(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Field string `json:"field"`
		From  string `json:"from"`
		To    string `json:"to"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	uuid, err := w.UnlinkField(payload.Field, payload.From, payload.To)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"uuid": uuid})
}
