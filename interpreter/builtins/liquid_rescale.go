package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func liquidRescale(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image    types.Image
		width    uint
		height   uint
		deltaX   = 1.0
		rigidity = 1.0
	)

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
	err := newImg.LiquidRescaleImage(width, height, deltaX, rigidity)
	if err != nil {
		return nil, fmt.Errorf("error liquid rescaling image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
