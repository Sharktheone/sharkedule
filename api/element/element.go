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
	return nil
}

// Detach Detaches an attachment from an element (delete from another element)
func Detach(c *fiber.Ctx) error {
	return nil
}

// Attach Attaches an element to another element (copy from another element)
func Attach(c *fiber.Ctx) error {
	return nil
}

// Attachments Gets all attachments of an element
func Attachments(c *fiber.Ctx) error {
	return nil
}

// List Lists all elements of a workspace / element (sub-elements)
func List(c *fiber.Ctx) error {
	return nil
}

// Create Creates a new element
func Create(c *fiber.Ctx) error {
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
