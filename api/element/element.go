package element

import (
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/gofiber/fiber/v2"
)

// Info Gets all properties (fields, attachments, etc) of an element
func Info(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	e, err := ea.ElementGet()
	if err != nil {
		return err
	}

	return c.JSON(e)

}

// Delete Deletes an element completely
func Delete(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	return ea.ElementDelete()
}

// Detach Detaches an attachment from an element (delete from another element)
func Detach(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Detach string `json:"detach"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementDetach(payload.Detach)
}

// Attach Attaches an element to another element (copy from another element)
func Attach(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Attach string `json:"attach"`
	})
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementAttach(payload.Attach)
}

// Attachments Gets all attachments of an element
func Attachments(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	return c.JSON(ea.ElementGetAttachments())
}

// Create Creates a new element
func Create(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
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

	_ = ea.ElementGetWorkspace()

	//TODO: parse Type and Fields

	return nil
}

// Update Updates an element (all properties)
func Update(c *fiber.Ctx) error {
	return nil
}

// Move Moves an element to another element (higher level API => could be done with attach/detach)
func Move(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Reference string `json:"reference"`
		To        string `json:"to"`
		Index     int    `json:"index,omitempty"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementMove(payload.Reference, payload.To, payload.Index)
}

// Copy Copies an element
func Copy(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		To    string `json:"to"`
		Index int    `json:"index,omitempty"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementCopy(payload.To, payload.Index)
}

func MoveElement(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Element string `json:"element"`
		To      string `json:"to"`
		Index   int    `json:"index,omitempty"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementMoveElement(payload.Element, payload.To, payload.Index)
}

func CopyElement(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Element string `json:"element"`
		To      string `json:"to"`
		Index   int    `json:"index,omitempty"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	return ea.ElementCopyElement(payload.Element, payload.To, payload.Index)
}

// GetType Gets the type of an element
func GetType(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"type": ea.ElementGetType()})
}

// UpdateType Updates the type of an element
func UpdateType(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Type string `json:"type"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	ty, err := types.ElementTypeFromString(payload.Type)
	if err != nil {
		return err
	}
	return ea.ElementUpdateType(&ty)
}

// List Lists all elements of a workspace / element (sub-elements)
func List(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	return c.JSON(ea.ElementGetSubElementsUUID())
}

// ListType -> Lists all elements of a workspace / element (sub-elements) of a specific type
func ListType(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Type string `json:"type"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	ty, err := types.ElementTypeFromString(payload.Type)
	if err != nil {
		return err
	}

	return c.JSON(ea.ElementGetSubElementsTypeUUID(ty))
}

// RecList -> Lists all elements of a workspace / element (sub-elements) recursively
func RecList(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	return c.JSON(ea.ElementGetRecSubElementsUUID())
}

// RecListType -> Lists all elements of a workspace / element (sub-elements) of a specific type recursively
func RecListType(c *fiber.Ctx) error {
	ea, err := middleware.ExtractElementAccess(c)
	if err != nil {
		return err
	}

	payload := new(struct {
		Type string `json:"type"`
	})

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	ty, err := types.ElementTypeFromString(payload.Type)
	if err != nil {
		return err
	}

	return c.JSON(ea.ElementGetRecSubElementsTypeUUID(ty))
}
