package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
	"github.com/fogo-sh/sorik/interpreter/types/enums"
)

func evaluateChannel(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image    types.Image
		channel  enums.ChannelType
		operator enums.EvaluateOperator
		value    float64
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"channel", &channel,
		"operator", &operator,
		"value", &value,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.EvaluateImageChannel(channel.Value, operator.Value, value)
	if err != nil {
		return nil, fmt.Errorf("error evaluating image channel: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
