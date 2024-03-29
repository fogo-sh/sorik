package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func rotate(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
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

	bgWand := imagick.NewPixelWand()
	bgWand.SetAlpha(0)

	err := newImg.RotateImage(bgWand, degrees)
	if err != nil {
		return nil, fmt.Errorf("error rotating image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
