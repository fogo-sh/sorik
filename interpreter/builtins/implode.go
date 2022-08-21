package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func implode(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image  types.Image
		radius float64
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"radius", &radius,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.ImplodeImage(radius)
	if err != nil {
		return nil, fmt.Errorf("error imploding image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
