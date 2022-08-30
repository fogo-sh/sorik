package builtins

import (
	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter/types/enums"
)

func ConstructBuiltins() starlark.StringDict {
	return starlark.StringDict{
		// Utilities
		"get_arg":    starlark.NewBuiltin("get_arg", getArg),
		"load_image": starlark.NewBuiltin("load_image", loadImage),

		// ImageMagick wrappers
		"liquid_rescale":   starlark.NewBuiltin("liquid_rescale", liquidRescale),
		"swirl":            starlark.NewBuiltin("swirl", swirl),
		"edge_detect":      starlark.NewBuiltin("edge_detect", edgeDetect),
		"invert":           starlark.NewBuiltin("invert", invert),
		"transform":        starlark.NewBuiltin("transform", transform),
		"contrast_stretch": starlark.NewBuiltin("contrast_stretch", contrastStretch),
		"sharpen":          starlark.NewBuiltin("sharpen", sharpen),
		"implode":          starlark.NewBuiltin("implode", implode),
		"modulate":         starlark.NewBuiltin("modulate", modulate),
		"gaussian_blur":    starlark.NewBuiltin("gaussian_blur", gaussianBlur),
		"flip":             starlark.NewBuiltin("flip", flip),
		"resize":           starlark.NewBuiltin("resize", resize),
		"evaluate_channel": starlark.NewBuiltin("evaluate_channel", evaluateChannel),
		"quantize":         starlark.NewBuiltin("quantize", quantize),
		"composite":        starlark.NewBuiltin("composite", composite),
		"rotate":           starlark.NewBuiltin("rotate", rotate),

		// Enums
		"FilterType":        enums.FilterTypeEnum{},
		"ChannelType":       enums.ChannelTypeEnum{},
		"EvaluateOperator":  enums.EvaluateOperatorEnum{},
		"ColorspaceType":    enums.ColorspaceTypeEnum{},
		"CompositeOperator": enums.CompositeOperatorEnum{},
	}
}
