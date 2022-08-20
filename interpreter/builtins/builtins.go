package builtins

import (
	"go.starlark.net/starlark"
)

func ConstructBuiltins() starlark.StringDict {
	return starlark.StringDict{
		// Utilities
		"get_arg":    starlark.NewBuiltin("get_arg", getArg),
		"load_image": starlark.NewBuiltin("load_image", loadImage),

		// ImageMagick wrappers
		"liquid_rescale": starlark.NewBuiltin("liquid_rescale", liquidRescale),
		"swirl":          starlark.NewBuiltin("swirl", swirl),
	}
}
