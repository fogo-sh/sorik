package types

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type ResizeFilter struct {
	FilterType imagick.FilterType
}

func (r ResizeFilter) String() string {
	//TODO implement me
	return fmt.Sprintf("ResizeFilter %d", r.FilterType)
}

func (r ResizeFilter) Type() string {
	return "ResizeFilter"
}

func (r ResizeFilter) Freeze() {
	return
}

func (r ResizeFilter) Truth() starlark.Bool {
	return true
}

func (r ResizeFilter) Hash() (uint32, error) {
	return uint32(r.FilterType), nil
}

var _ starlark.Value = (*ResizeFilter)(nil)
