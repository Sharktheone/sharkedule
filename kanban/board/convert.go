package board

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sharkedule/database/types"
)

func (b *Board) Convert() (*types.Board, error) {
	var board *types.Board
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &board,
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

func ConvertBoard(board *types.Board) (*Board, error) {
	var b *Board
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &b,
		TagName: "json", // TODO: Change to other tag => to also convert index etc.
	})
	if err != nil {
		return nil, fmt.Errorf("failed creating board decoder: %v", err)
	}
	if err := decoder.Decode(board); err != nil {
		return nil, fmt.Errorf("failed decoding board: %v", err)
	}
	return b, nil
}

func ConvertBoards(boards []*types.Board) ([]*Board, error) {
	var b []*Board
	for _, board := range boards {
		converted, err := ConvertBoard(board)
		if err != nil {
			return nil, err
		}
		b = append(b, converted)
	}
	return b, nil
}
