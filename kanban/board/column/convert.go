package column

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sharkedule/database/types"
)

func (c *Column) Convert() (*types.Column, error) {
	var column *types.Column
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &column,
		TagName: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("failed creating column decoder: %v", err)
	}
	if err := decoder.Decode(c); err != nil {
		return nil, fmt.Errorf("failed decoding column: %v", err)
	}
	return column, nil
}
