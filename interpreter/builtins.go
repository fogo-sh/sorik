package interpreter

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func loadImage() *starlark.Builtin {
	return starlark.NewBuiltin(
		"load_image",
		func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
			var url string
			if err := starlark.UnpackArgs(fn.Name(), args, kwargs, "url", &url); err != nil {
				return nil, err
			}

			log.Debug().Str("url", url).Msg("Downloading image")
			resp, err := http.Get(url)
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
			err = inputImg.ReadImageBlob(buffer.Bytes())
			if err != nil {
				return nil, fmt.Errorf("error reading image: %w", err)
			}

			return &Image{wand: inputImg}, nil
		},
	)
}

func constructBuiltins() starlark.StringDict {
	return starlark.StringDict{
		"load_image": loadImage(),
	}
}
