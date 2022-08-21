// Code generated by enum_generator; DO NOT EDIT.
//
//go:generate go run enum_generator.go FilterType filter_type.go
package types

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type FilterType struct {
	Value     imagick.FilterType
	StringVal string
}

func (v FilterType) String() string {
	return fmt.Sprintf("FilterType %s", v.StringVal)
}

func (v FilterType) Type() string {
	return "%s"
}

func (v FilterType) Freeze() {
	return
}

func (v FilterType) Truth() starlark.Bool {
	return true
}

func (v FilterType) Hash() (uint32, error) {
	return uint32(v.Value), nil
}

var _ starlark.Value = (*FilterType)(nil)

var _FilterTypeMap = map[string]FilterType{
	"FILTER_UNDEFINED":      {imagick.FILTER_UNDEFINED, "FILTER_UNDEFINED"},
	"FILTER_POINT":          {imagick.FILTER_POINT, "FILTER_POINT"},
	"FILTER_BOX":            {imagick.FILTER_BOX, "FILTER_BOX"},
	"FILTER_TRIANGLE":       {imagick.FILTER_TRIANGLE, "FILTER_TRIANGLE"},
	"FILTER_HERMITE":        {imagick.FILTER_HERMITE, "FILTER_HERMITE"},
	"FILTER_HANNING":        {imagick.FILTER_HANNING, "FILTER_HANNING"},
	"FILTER_HAMMING":        {imagick.FILTER_HAMMING, "FILTER_HAMMING"},
	"FILTER_BLACKMAN":       {imagick.FILTER_BLACKMAN, "FILTER_BLACKMAN"},
	"FILTER_GAUSSIAN":       {imagick.FILTER_GAUSSIAN, "FILTER_GAUSSIAN"},
	"FILTER_QUADRATIC":      {imagick.FILTER_QUADRATIC, "FILTER_QUADRATIC"},
	"FILTER_CUBIC":          {imagick.FILTER_CUBIC, "FILTER_CUBIC"},
	"FILTER_CATROM":         {imagick.FILTER_CATROM, "FILTER_CATROM"},
	"FILTER_MITCHELL":       {imagick.FILTER_MITCHELL, "FILTER_MITCHELL"},
	"FILTER_JINC":           {imagick.FILTER_JINC, "FILTER_JINC"},
	"FILTER_SINC":           {imagick.FILTER_SINC, "FILTER_SINC"},
	"FILTER_SINC_FAST":      {imagick.FILTER_SINC_FAST, "FILTER_SINC_FAST"},
	"FILTER_KAISER":         {imagick.FILTER_KAISER, "FILTER_KAISER"},
	"FILTER_WELSH":          {imagick.FILTER_WELSH, "FILTER_WELSH"},
	"FILTER_PARZEN":         {imagick.FILTER_PARZEN, "FILTER_PARZEN"},
	"FILTER_BOHMAN":         {imagick.FILTER_BOHMAN, "FILTER_BOHMAN"},
	"FILTER_BARTLETT":       {imagick.FILTER_BARTLETT, "FILTER_BARTLETT"},
	"FILTER_LAGRANGE":       {imagick.FILTER_LAGRANGE, "FILTER_LAGRANGE"},
	"FILTER_LANCZOS":        {imagick.FILTER_LANCZOS, "FILTER_LANCZOS"},
	"FILTER_LANCZOS_SHARP":  {imagick.FILTER_LANCZOS_SHARP, "FILTER_LANCZOS_SHARP"},
	"FILTER_LANCZOS2":       {imagick.FILTER_LANCZOS2, "FILTER_LANCZOS2"},
	"FILTER_LANCZOS2_SHARP": {imagick.FILTER_LANCZOS2_SHARP, "FILTER_LANCZOS2_SHARP"},
	"FILTER_ROBIDOUX":       {imagick.FILTER_ROBIDOUX, "FILTER_ROBIDOUX"},
	"FILTER_ROBIDOUX_SHARP": {imagick.FILTER_ROBIDOUX_SHARP, "FILTER_ROBIDOUX_SHARP"},
	"FILTER_COSINE":         {imagick.FILTER_COSINE, "FILTER_COSINE"},
	"FILTER_SPLINE":         {imagick.FILTER_SPLINE, "FILTER_SPLINE"},
	"FILTER_SENTINEL":       {imagick.FILTER_SENTINEL, "FILTER_SENTINEL"},
	"FILTER_LANCZOS_RADIUS": {imagick.FILTER_LANCZOS_RADIUS, "FILTER_LANCZOS_RADIUS"},
}

type FilterTypeEnum struct{}

func (v FilterTypeEnum) String() string {
	return "FilterTypeEnum"
}

func (v FilterTypeEnum) Type() string {
	return "FilterTypeEnum"
}

func (v FilterTypeEnum) Freeze() {
	return
}

func (v FilterTypeEnum) Truth() starlark.Bool {
	return true
}

func (v FilterTypeEnum) Hash() (uint32, error) {
	return 0, nil
}

func (v FilterTypeEnum) Attr(name string) (starlark.Value, error) {
	val, found := _FilterTypeMap[name]
	if !found {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("unknown FilterType %s", name))
	}
	return val, nil
}

func (v FilterTypeEnum) AttrNames() []string {
	var attrNames []string

	for name := range _FilterTypeMap {
		attrNames = append(attrNames, name)
	}

	return attrNames
}

var _ starlark.Value = (*FilterTypeEnum)(nil)
var _ starlark.HasAttrs = (*FilterTypeEnum)(nil)
