package element

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
)

type Element interface {
	GetUUID() string
	GetName() string
	GetType() elements.Type
	GetFields() []field.Field
}
