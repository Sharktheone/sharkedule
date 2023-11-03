package column

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
)

type Column struct {
	*types.Column
	Workspace string
}
