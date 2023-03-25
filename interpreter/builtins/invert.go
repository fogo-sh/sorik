package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func invert(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image         types.Image
		greyscaleOnly = false
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"greyscale_only?", &greyscaleOnly,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()

	origMask := newImg.SetImageChannelMask(imagick.CHANNEL_RED | imagick.CHANNEL_GREEN | imagick.CHANNEL_BLUE)

	err := newImg.NegateImage(greyscaleOnly)
	if err != nil {
		return nil, fmt.Errorf("error inverting image: %w", err)
	}

	newImg.SetImageChannelMask(origMask)

	return types.Image{Wand: newImg}, nil
}
