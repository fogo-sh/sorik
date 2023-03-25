package builtins

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/rs/zerolog/log"
	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v3/imagick"

	"github.com/fogo-sh/sorik/interpreter/types"
)

func loadImage(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		imageUrl string
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"url", &imageUrl,
	); err != nil {
		return nil, err
	}

	log.Debug().Str("url", imageUrl).Msg("Downloading image")
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, fmt.Errorf("error downloading image: %w", err)
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)

	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error copying image to buffer: %w", err)
	}

	inputImg := imagick.NewMagickWand()

	parsedUrl, _ := url.Parse(imageUrl)
	filename := path.Base(parsedUrl.Path)
	err = inputImg.SetFilename(filename)
	if err != nil {
		log.Error().Err(err).Msg("Failed to set image filename - loading may not behave as expected.")
	}

	err = inputImg.ReadImageBlob(buffer.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error reading image: %w", err)
	}

	return types.Image{Wand: inputImg}, nil
}
