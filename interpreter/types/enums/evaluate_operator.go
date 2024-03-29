// Code generated by enum_generator; DO NOT EDIT.
//
//go:generate go run enum_generator.go
package enums

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type EvaluateOperator struct {
	Value     imagick.EvaluateOperator
	stringVal string
}

func (v EvaluateOperator) String() string {
	return fmt.Sprintf("EvaluateOperator %s", v.stringVal)
}

func (v EvaluateOperator) Type() string {
	return "EvaluateOperator"
}

func (v EvaluateOperator) Freeze() {
	return
}

func (v EvaluateOperator) Truth() starlark.Bool {
	return true
}

func (v EvaluateOperator) Hash() (uint32, error) {
	return uint32(v.Value), nil
}

var _ starlark.Value = (*EvaluateOperator)(nil)

var _EvaluateOperatorMap = map[string]EvaluateOperator{
	"EVAL_OP_UNDEFINED":            {imagick.EVAL_OP_UNDEFINED, "EVAL_OP_UNDEFINED"},
	"EVAL_OP_ADD":                  {imagick.EVAL_OP_ADD, "EVAL_OP_ADD"},
	"EVAL_OP_AND":                  {imagick.EVAL_OP_AND, "EVAL_OP_AND"},
	"EVAL_OP_DIVIDE":               {imagick.EVAL_OP_DIVIDE, "EVAL_OP_DIVIDE"},
	"EVAL_OP_LEFT_SHIFT":           {imagick.EVAL_OP_LEFT_SHIFT, "EVAL_OP_LEFT_SHIFT"},
	"EVAL_OP_MAX":                  {imagick.EVAL_OP_MAX, "EVAL_OP_MAX"},
	"EVAL_OP_MIN":                  {imagick.EVAL_OP_MIN, "EVAL_OP_MIN"},
	"EVAL_OP_MULTIPLY":             {imagick.EVAL_OP_MULTIPLY, "EVAL_OP_MULTIPLY"},
	"EVAL_OP_OR":                   {imagick.EVAL_OP_OR, "EVAL_OP_OR"},
	"EVAL_OP_RIGHT_SHIFT":          {imagick.EVAL_OP_RIGHT_SHIFT, "EVAL_OP_RIGHT_SHIFT"},
	"EVAL_OP_SET":                  {imagick.EVAL_OP_SET, "EVAL_OP_SET"},
	"EVAL_OP_SUBTRACT":             {imagick.EVAL_OP_SUBTRACT, "EVAL_OP_SUBTRACT"},
	"EVAL_OP_XOR":                  {imagick.EVAL_OP_XOR, "EVAL_OP_XOR"},
	"EVAL_OP_POW":                  {imagick.EVAL_OP_POW, "EVAL_OP_POW"},
	"EVAL_OP_LOG":                  {imagick.EVAL_OP_LOG, "EVAL_OP_LOG"},
	"EVAL_OP_THRESHOLD":            {imagick.EVAL_OP_THRESHOLD, "EVAL_OP_THRESHOLD"},
	"EVAL_OP_THRESHOLD_BLACK":      {imagick.EVAL_OP_THRESHOLD_BLACK, "EVAL_OP_THRESHOLD_BLACK"},
	"EVAL_OP_THRESHOLD_WHITE":      {imagick.EVAL_OP_THRESHOLD_WHITE, "EVAL_OP_THRESHOLD_WHITE"},
	"EVAL_OP_GAUSSIAN_NOISE":       {imagick.EVAL_OP_GAUSSIAN_NOISE, "EVAL_OP_GAUSSIAN_NOISE"},
	"EVAL_OP_IMPULSE_NOISE":        {imagick.EVAL_OP_IMPULSE_NOISE, "EVAL_OP_IMPULSE_NOISE"},
	"EVAL_OP_LAPLACIAN_NOISE":      {imagick.EVAL_OP_LAPLACIAN_NOISE, "EVAL_OP_LAPLACIAN_NOISE"},
	"EVAL_OP_MULTIPLICATIVE_NOISE": {imagick.EVAL_OP_MULTIPLICATIVE_NOISE, "EVAL_OP_MULTIPLICATIVE_NOISE"},
	"EVAL_OP_POISSON_NOISE":        {imagick.EVAL_OP_POISSON_NOISE, "EVAL_OP_POISSON_NOISE"},
	"EVAL_OP_UNIFORM_NOISE":        {imagick.EVAL_OP_UNIFORM_NOISE, "EVAL_OP_UNIFORM_NOISE"},
	"EVAL_OP_COSINE":               {imagick.EVAL_OP_COSINE, "EVAL_OP_COSINE"},
	"EVAL_OP_SINE":                 {imagick.EVAL_OP_SINE, "EVAL_OP_SINE"},
	"EVAL_OP_ADD_MODULUS":          {imagick.EVAL_OP_ADD_MODULUS, "EVAL_OP_ADD_MODULUS"},
	"EVAL_OP_MEAN":                 {imagick.EVAL_OP_MEAN, "EVAL_OP_MEAN"},
	"EVAL_OP_ABS":                  {imagick.EVAL_OP_ABS, "EVAL_OP_ABS"},
	"EVAL_OP_EXPONENTIAL":          {imagick.EVAL_OP_EXPONENTIAL, "EVAL_OP_EXPONENTIAL"},
	"EVAL_OP_MEDIAN":               {imagick.EVAL_OP_MEDIAN, "EVAL_OP_MEDIAN"},
	"EVAL_OP_SUM":                  {imagick.EVAL_OP_SUM, "EVAL_OP_SUM"},
}

var _EvaluateOperatorNames = []string{
	"EVAL_OP_UNDEFINED",
	"EVAL_OP_ADD",
	"EVAL_OP_AND",
	"EVAL_OP_DIVIDE",
	"EVAL_OP_LEFT_SHIFT",
	"EVAL_OP_MAX",
	"EVAL_OP_MIN",
	"EVAL_OP_MULTIPLY",
	"EVAL_OP_OR",
	"EVAL_OP_RIGHT_SHIFT",
	"EVAL_OP_SET",
	"EVAL_OP_SUBTRACT",
	"EVAL_OP_XOR",
	"EVAL_OP_POW",
	"EVAL_OP_LOG",
	"EVAL_OP_THRESHOLD",
	"EVAL_OP_THRESHOLD_BLACK",
	"EVAL_OP_THRESHOLD_WHITE",
	"EVAL_OP_GAUSSIAN_NOISE",
	"EVAL_OP_IMPULSE_NOISE",
	"EVAL_OP_LAPLACIAN_NOISE",
	"EVAL_OP_MULTIPLICATIVE_NOISE",
	"EVAL_OP_POISSON_NOISE",
	"EVAL_OP_UNIFORM_NOISE",
	"EVAL_OP_COSINE",
	"EVAL_OP_SINE",
	"EVAL_OP_ADD_MODULUS",
	"EVAL_OP_MEAN",
	"EVAL_OP_ABS",
	"EVAL_OP_EXPONENTIAL",
	"EVAL_OP_MEDIAN",
	"EVAL_OP_SUM",
}

type EvaluateOperatorEnum struct{}

func (v EvaluateOperatorEnum) String() string {
	return "EvaluateOperatorEnum"
}

func (v EvaluateOperatorEnum) Type() string {
	return "EvaluateOperatorEnum"
}

func (v EvaluateOperatorEnum) Freeze() {
	return
}

func (v EvaluateOperatorEnum) Truth() starlark.Bool {
	return true
}

func (v EvaluateOperatorEnum) Hash() (uint32, error) {
	return 0, nil
}

func (v EvaluateOperatorEnum) Attr(name string) (starlark.Value, error) {
	val, found := _EvaluateOperatorMap[name]
	if !found {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("unknown EvaluateOperator %s", name))
	}
	return val, nil
}

func (v EvaluateOperatorEnum) AttrNames() []string {
	return _EvaluateOperatorNames
}

var _ starlark.Value = (*EvaluateOperatorEnum)(nil)
var _ starlark.HasAttrs = (*EvaluateOperatorEnum)(nil)
