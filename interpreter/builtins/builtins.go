package builtins

import (
	"go.starlark.net/starlark"
)

func ConstructBuiltins() starlark.StringDict {
	return starlark.StringDict{
		"load_image":     starlark.NewBuiltin("load_image", loadImage),
		"liquid_rescale": starlark.NewBuiltin("liquid_rescale", liquidRescale),
	}
}
