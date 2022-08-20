package builtins

import (
	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func transform(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var crop = starlark.String("")
	var geometry = starlark.String("")

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"crop?", &crop,
		"geometry?", &geometry,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.TransformImage(string(crop), string(geometry))

	return types.Image{Wand: newImg}, nil
}
