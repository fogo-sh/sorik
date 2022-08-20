package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func swirl(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var degrees float64

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"degrees", &degrees,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.SwirlImage(degrees)
	if err != nil {
		return nil, fmt.Errorf("error swirling image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
