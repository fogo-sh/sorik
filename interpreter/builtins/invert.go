package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func invert(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var greyscaleOnly = starlark.False

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"greyscale_only?", &greyscaleOnly,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.NegateImage(bool(greyscaleOnly))
	if err != nil {
		return nil, fmt.Errorf("error inverting image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
