package element

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
)

type Element interface {
	UUID() string
	Name() string
	Type() elements.Type
	Fields() []field.Field
}
