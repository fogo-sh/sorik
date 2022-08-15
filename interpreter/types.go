package interpreter

import (
	"errors"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type Image struct {
	wand *imagick.MagickWand
}

func (i *Image) String() string {
	//TODO add better string representation
	return i.wand.IdentifyImage()
}

func (i *Image) Type() string {
	return "Image"
}

func (i *Image) Freeze() {
	//TODO implement me
	return
}

func (i *Image) Truth() starlark.Bool {
	return starlark.True
}

func (i *Image) Hash() (uint32, error) {
	return 0, errors.New("not implemented")
}

var _ starlark.Value = (*Image)(nil)
