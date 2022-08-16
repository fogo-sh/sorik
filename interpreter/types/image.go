package types

import (
	"errors"
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type Image struct {
	Wand *imagick.MagickWand
}

func (i *Image) String() string {
	//TODO add better string representation
	return i.Wand.IdentifyImage()
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
	//TODO implement me
	return 0, errors.New("not implemented")
}

func (i *Image) AttrNames() []string {
	return []string{"width", "height"}
}

func (i *Image) Attr(name string) (starlark.Value, error) {
	switch name {
	case "height":
		return starlark.MakeInt(int(i.Wand.GetImageHeight())), nil
	case "width":
		return starlark.MakeInt(int(i.Wand.GetImageWidth())), nil
	default:
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("type image has no attribute %s", name))
	}
}

var _ starlark.Value = (*Image)(nil)
var _ starlark.HasAttrs = (*Image)(nil)
