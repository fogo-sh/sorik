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
		"composite":        starlark.NewBuiltin("composite", composite),
		"contrast_stretch": starlark.NewBuiltin("contrast_stretch", contrastStretch),
		"edge_detect":      starlark.NewBuiltin("edge_detect", edgeDetect),
		"evaluate_channel": starlark.NewBuiltin("evaluate_channel", evaluateChannel),
		"flip":             starlark.NewBuiltin("flip", flip),
		"gaussian_blur":    starlark.NewBuiltin("gaussian_blur", gaussianBlur),
		"implode":          starlark.NewBuiltin("implode", implode),
		"invert":           starlark.NewBuiltin("invert", invert),
		"liquid_rescale":   starlark.NewBuiltin("liquid_rescale", liquidRescale),
		"modulate":         starlark.NewBuiltin("modulate", modulate),
		"quantize":         starlark.NewBuiltin("quantize", quantize),
		"resize":           starlark.NewBuiltin("resize", resize),
		"rotate":           starlark.NewBuiltin("rotate", rotate),
		"sharpen":          starlark.NewBuiltin("sharpen", sharpen),
		"swirl":            starlark.NewBuiltin("swirl", swirl),
		"transform":        starlark.NewBuiltin("transform", transform),

		// Enums
		"ChannelType":       enums.ChannelTypeEnum{},
		"ColorspaceType":    enums.ColorspaceTypeEnum{},
		"CompositeOperator": enums.CompositeOperatorEnum{},
		"EvaluateOperator":  enums.EvaluateOperatorEnum{},
		"FilterType":        enums.FilterTypeEnum{},
	}
}
