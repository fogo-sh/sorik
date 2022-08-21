// Code generated by enum_generator; DO NOT EDIT.
//
//go:generate go run enum_generator.go ChannelType channel_type.go
package types

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type ChannelType struct {
	Value     imagick.ChannelType
	StringVal string
}

func (v ChannelType) String() string {
	return fmt.Sprintf("ChannelType %s", v.StringVal)
}

func (v ChannelType) Type() string {
	return "%s"
}

func (v ChannelType) Freeze() {
	return
}

func (v ChannelType) Truth() starlark.Bool {
	return true
}

func (v ChannelType) Hash() (uint32, error) {
	return uint32(v.Value), nil
}

var _ starlark.Value = (*ChannelType)(nil)

var _ChannelTypeMap = map[string]ChannelType{
	"CHANNEL_UNDEFINED":  {imagick.CHANNEL_UNDEFINED, "CHANNEL_UNDEFINED"},
	"CHANNEL_RED":        {imagick.CHANNEL_RED, "CHANNEL_RED"},
	"CHANNEL_GRAY":       {imagick.CHANNEL_GRAY, "CHANNEL_GRAY"},
	"CHANNEL_CYAN":       {imagick.CHANNEL_CYAN, "CHANNEL_CYAN"},
	"CHANNEL_GREEN":      {imagick.CHANNEL_GREEN, "CHANNEL_GREEN"},
	"CHANNEL_MAGENTA":    {imagick.CHANNEL_MAGENTA, "CHANNEL_MAGENTA"},
	"CHANNEL_BLUE":       {imagick.CHANNEL_BLUE, "CHANNEL_BLUE"},
	"CHANNEL_YELLOW":     {imagick.CHANNEL_YELLOW, "CHANNEL_YELLOW"},
	"CHANNEL_ALPHA":      {imagick.CHANNEL_ALPHA, "CHANNEL_ALPHA"},
	"CHANNEL_OPACITY":    {imagick.CHANNEL_OPACITY, "CHANNEL_OPACITY"},
	"CHANNEL_BLACK":      {imagick.CHANNEL_BLACK, "CHANNEL_BLACK"},
	"CHANNEL_INDEX":      {imagick.CHANNEL_INDEX, "CHANNEL_INDEX"},
	"CHANNEL_TRUE_ALPHA": {imagick.CHANNEL_TRUE_ALPHA, "CHANNEL_TRUE_ALPHA"},
	"CHANNELS_COMPOSITE": {imagick.CHANNELS_COMPOSITE, "CHANNELS_COMPOSITE"},
	"CHANNELS_ALL":       {imagick.CHANNELS_ALL, "CHANNELS_ALL"},
	"CHANNELS_RGB":       {imagick.CHANNELS_RGB, "CHANNELS_RGB"},
	"CHANNELS_GRAY":      {imagick.CHANNELS_GRAY, "CHANNELS_GRAY"},
	"CHANNELS_SYNC":      {imagick.CHANNELS_SYNC, "CHANNELS_SYNC"},
	"CHANNELS_DEFAULT":   {imagick.CHANNELS_DEFAULT, "CHANNELS_DEFAULT"},
}

type ChannelTypeEnum struct{}

func (v ChannelTypeEnum) String() string {
	return "ChannelTypeEnum"
}

func (v ChannelTypeEnum) Type() string {
	return "ChannelTypeEnum"
}

func (v ChannelTypeEnum) Freeze() {
	return
}

func (v ChannelTypeEnum) Truth() starlark.Bool {
	return true
}

func (v ChannelTypeEnum) Hash() (uint32, error) {
	return 0, nil
}

func (v ChannelTypeEnum) Attr(name string) (starlark.Value, error) {
	val, found := _ChannelTypeMap[name]
	if !found {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("unknown ChannelType %s", name))
	}
	return val, nil
}

func (v ChannelTypeEnum) AttrNames() []string {
	var attrNames []string

	for name := range _ChannelTypeMap {
		attrNames = append(attrNames, name)
	}

	return attrNames
}

var _ starlark.Value = (*ChannelTypeEnum)(nil)
var _ starlark.HasAttrs = (*ChannelTypeEnum)(nil)
