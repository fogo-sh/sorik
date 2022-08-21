package builtins

import (
	"errors"
	"fmt"

	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func flip(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		image      types.Image
		horizontal = false
		vertical   = false
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"image", &image,
		"horizontal?", &horizontal,
		"vertical?", &vertical,
	); err != nil {
		return nil, err
	}

	if !horizontal && !vertical {
		return nil, errors.New("no flip direction provided")
	}

	newImg := image.Wand.Clone()

	if horizontal {
		err := newImg.FlopImage()
		if err != nil {
			return nil, fmt.Errorf("error flipping image: %w", err)
		}
	}
	if vertical {
		err := newImg.FlipImage()
		if err != nil {
			return nil, fmt.Errorf("error flipping image: %w", err)
		}
	}

	return types.Image{Wand: newImg}, nil
}
