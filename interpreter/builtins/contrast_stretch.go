package builtins

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func contrastStretch(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var image types.Image
	var blackPoint float64
	var whitePoint float64

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"black_point", &blackPoint,
		"white_point", &whitePoint,
	); err != nil {
		return nil, err
	}

	newImg := image.Wand.Clone()
	err := newImg.ContrastStretchImage(blackPoint, whitePoint)
	if err != nil {
		return nil, fmt.Errorf("error contrast stretching image: %w", err)
	}

	return types.Image{Wand: newImg}, nil
}
