package board

import (
	"github.com/google/uuid"
)

func NewBoard(name string) *Board {
	return &Board{
		UUID: uuid.New().String(),
		Name: name,
	}
}

type A struct {
	Items []AItem
}

type AItem struct {
	Name string
}

func (a *A) AddItem(name string) {
	a.Items = append(a.Items, AItem{Name: name})
}

type B struct {
	Items []*BItem
}

type BItem struct {
	Name string
}
