package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func composite(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		targetImage types.Image
		sourceImage types.Image
		operator    types.CompositeOperator
		x           int
		y           int
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"target_image", &targetImage,
		"source_image", &sourceImage,
		"operator", &operator,
		"x", &x,
		"y", &y,
	); err != nil {
		return nil, err
	}

	newImg := targetImage.Wand.Clone()
	err := newImg.CompositeImage(sourceImage.Wand, operator.Value, x, y)
	if err != nil {
		return nil, fmt.Errorf("error compositing image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
