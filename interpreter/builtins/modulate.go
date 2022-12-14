package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func modulate(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image      types.Image
		brightness = 100.0
		saturation = 100.0
		hue        = 100.0
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"brightness?", &brightness,
		"saturation?", &saturation,
		"hue?", &hue,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.ModulateImage(brightness, saturation, hue)
	if err != nil {
		return nil, fmt.Errorf("error modulating image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
