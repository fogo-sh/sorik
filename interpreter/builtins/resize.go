package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func resize(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image  types.Image
		width  uint
		height uint
		filter types.ResizeFilter
		blur   float64
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"width", &width,
		"height", &height,
		"filter", &filter,
		"blur", &blur,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.ResizeImage(width, height, filter.FilterType, blur)
	if err != nil {
		return nil, fmt.Errorf("error swirling image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
