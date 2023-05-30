package board

import (
	"github.com/mitchellh/mapstructure"
	"sharkedule/database/types"
)

func (b *Board) Convert() (*types.Board, error) {
	var board *types.Board
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  board,
		TagName: "json",
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(b); err != nil {
		return nil, err
	}
	return board, nil
}
