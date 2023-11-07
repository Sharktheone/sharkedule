package tag

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
)

type Tag struct {
	*types.Tag
	Workspace string
}
