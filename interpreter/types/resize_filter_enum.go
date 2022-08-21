package types

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

var resizeFilterMap = map[string]ResizeFilter{
	"Undefined":     {imagick.FILTER_UNDEFINED},
	"Point":         {imagick.FILTER_POINT},
	"Box":           {imagick.FILTER_BOX},
	"Triangle":      {imagick.FILTER_TRIANGLE},
	"Hermite":       {imagick.FILTER_HERMITE},
	"Hanning":       {imagick.FILTER_HANNING},
	"Hamming":       {imagick.FILTER_HAMMING},
	"Blackman":      {imagick.FILTER_BLACKMAN},
	"Gaussian":      {imagick.FILTER_GAUSSIAN},
	"Quadratic":     {imagick.FILTER_QUADRATIC},
	"Cubic":         {imagick.FILTER_CUBIC},
	"Catrom":        {imagick.FILTER_CATROM},
	"Mitchell":      {imagick.FILTER_MITCHELL},
	"Jinc":          {imagick.FILTER_JINC},
	"Sinc":          {imagick.FILTER_SINC},
	"SincFast":      {imagick.FILTER_SINC_FAST},
	"Kaiser":        {imagick.FILTER_KAISER},
	"Welsh":         {imagick.FILTER_WELSH},
	"Parzen":        {imagick.FILTER_PARZEN},
	"Bohman":        {imagick.FILTER_BOHMAN},
	"Bartlett":      {imagick.FILTER_BARTLETT},
	"Lagrange":      {imagick.FILTER_LAGRANGE},
	"Lanczos":       {imagick.FILTER_LANCZOS},
	"LanczosSharp":  {imagick.FILTER_LANCZOS_SHARP},
	"Lanczos2":      {imagick.FILTER_LANCZOS2},
	"Lanczos2Sharp": {imagick.FILTER_LANCZOS2_SHARP},
	"Robidoux":      {imagick.FILTER_ROBIDOUX},
	"RobidouxSharp": {imagick.FILTER_ROBIDOUX_SHARP},
	"Cosine":        {imagick.FILTER_COSINE},
	"Spline":        {imagick.FILTER_SPLINE},
	"Sentinel":      {imagick.FILTER_SENTINEL},
	"LanczosRadius": {imagick.FILTER_LANCZOS_RADIUS},
}

type ResizeFilterEnum struct {
	FilterType imagick.FilterType
}

func (r ResizeFilterEnum) String() string {
	return "ResizeFilterEnum"
}

func (r ResizeFilterEnum) Type() string {
	return "ResizeFilterEnum"
}

func (r ResizeFilterEnum) Freeze() {
	return
}

func (r ResizeFilterEnum) Truth() starlark.Bool {
	return true
}

func (r ResizeFilterEnum) Hash() (uint32, error) {
	return uint32(r.FilterType), nil
}

func (r ResizeFilterEnum) Attr(name string) (starlark.Value, error) {
	val, found := resizeFilterMap[name]
	if !found {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("unknown ResizeFilter %s", name))
	}
	return val, nil
}

func (r ResizeFilterEnum) AttrNames() []string {
	var attrNames []string

	for name := range resizeFilterMap {
		attrNames = append(attrNames, name)
	}

	return attrNames
}

var _ starlark.Value = (*ResizeFilterEnum)(nil)
var _ starlark.HasAttrs = (*ResizeFilterEnum)(nil)
