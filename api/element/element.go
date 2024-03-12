package element

import (
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// Info Gets all properties (fields, attachments, etc) of an element
func Info(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	return c.JSON(e)

}

// Delete Deletes an element completely
func Delete(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	return e.Delete()
}

// Detach Detaches an attachment from an element (delete from another element)
func Detach(c *fiber.Ctx) error {
	_, _, err := middleware.ExtractElement(c)
	if err != nil {
		return err

	}
	return nil
}

// Attach Attaches an element to another element (copy from another element)
func Attach(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Attach string `json:"attach"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return e.Attach(payload.Attach)
}

// Attachments Gets all attachments of an element
func Attachments(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	return c.JSON(e.GetAttachments())
}

// List Lists all elements of a workspace / element (sub-elements) //TODO: Decide if this is not basically a duplicate of Attachments
func List(c *fiber.Ctx) error {
	return nil
}

// Create Creates a new element
func Create(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Type   string            `json:"type"`
		Name   string            `json:"name"`
		Fields map[string]string `json:"fields,omitempty"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	_ = e.GetWorkspace()

	//TODO: parse Type and Fields

	return nil
}

// Update Updates an element (all properties)
func Update(c *fiber.Ctx) error {
	return nil
}

// Move Moves an element to another element (higher level API => could be done with attach/detach)
func Move(c *fiber.Ctx) error {
	return nil
}

// Copy Copies an element
func Copy(c *fiber.Ctx) error {
	return nil
}

// GetType Gets the type of an element
func GetType(c *fiber.Ctx) error {
	return nil
}

// UpdateType Updates the type of an element
func UpdateType(c *fiber.Ctx) error {
	return nil
}

// ListType -> Lists all elements of a workspace / element (sub-elements) of a specific type
func ListType(c *fiber.Ctx) error {
	return nil
}
