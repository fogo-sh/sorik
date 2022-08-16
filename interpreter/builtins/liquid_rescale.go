package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func liquidRescale(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var width uint
	var height uint
	var deltaX = starlark.Float(1)
	var rigidity = starlark.Float(1)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"width", &width,
		"height", &height,
		"deltax?", &deltaX,
		"rigidity?", &rigidity,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.LiquidRescaleImage(width, height, float64(deltaX), float64(rigidity))
	if err != nil {
		return nil, fmt.Errorf("error liquid rescaling image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
