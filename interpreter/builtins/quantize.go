package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
	"github.com/fogo-sh/sorik/interpreter/types/enums"
)

func quantize(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image          types.Image
		numColors      uint
		colorspaceType enums.ColorspaceType
		treeDepth      uint
		dither         bool
		measureError   bool
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"num_colors", &numColors,
		"colorspace_type", &colorspaceType,
		"tree_depth", &treeDepth,
		"dither", &dither,
		"measure_error", &measureError,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.QuantizeImage(numColors, colorspaceType.Value, treeDepth, imagick.DITHER_METHOD_FLOYD_STEINBERG, dither)
	if err != nil {
		return nil, fmt.Errorf("error quantizing image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
