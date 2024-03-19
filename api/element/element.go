package element

import (
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/Sharktheone/sharkedule/types"
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
	_, e, err := middleware.ExtractElement(c)
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

	return e.Move(payload.Reference, payload.To, payload.Index)
}

// Copy Copies an element
func Copy(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
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

	return e.Copy(payload.To, payload.Index)
}

func MoveElement(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
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

	return e.MoveElement(payload.Element, payload.To, payload.Index)
}

func CopyElement(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
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

	return e.CopyElement(payload.Element, payload.To, payload.Index)
}

// GetType Gets the type of an element
func GetType(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	return c.JSON(api.JSON{"type": e.GetType()})
}

// UpdateType Updates the type of an element
func UpdateType(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
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
	return e.UpdateType(&ty)
}

// ListType -> Lists all elements of a workspace / element (sub-elements) of a specific type
func ListType(c *fiber.Ctx) error {
	_, e, err := middleware.ExtractElement(c)
	if err != nil {
		return err
	}

	sub := e.GetSubElements()

	type subType struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	}

	subTypes := make([]subType, len(sub))

	for i, e := range sub {
		subTypes[i] = subType{UUID: e.GetUUID(), Type: string(*e.GetType())}
	}

	return c.JSON(subTypes)
}
