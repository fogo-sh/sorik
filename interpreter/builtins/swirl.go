package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func swirl(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image   types.Image
		degrees float64
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"degrees", &degrees,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.SwirlImage(degrees, imagick.INTERPOLATE_PIXEL_BILINEAR)
	if err != nil {
		return nil, fmt.Errorf("error swirling image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
