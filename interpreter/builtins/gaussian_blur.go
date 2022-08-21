package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func gaussianBlur(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var radius float64
	var sigma float64

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"radius", &radius,
		"sigma", &sigma,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.GaussianBlurImage(radius, sigma)
	if err != nil {
		return nil, fmt.Errorf("error gaussian blurring image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
