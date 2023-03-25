package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func edgeDetect(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image  types.Image
		radius float64
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"radius", &radius,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()

	origMask := newImg.SetImageChannelMask(imagick.CHANNEL_RED | imagick.CHANNEL_GREEN | imagick.CHANNEL_BLUE)

	err := newImg.EdgeImage(radius)
	if err != nil {
		return nil, fmt.Errorf("error edge detecting image: %w", err)
	}

	newImg.SetImageChannelMask(origMask)

	return types.Image{Wand: newImg}, nil
}
