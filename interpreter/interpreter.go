package interpreter

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"

	"github.com/fogo-sh/sorik/interpreter/builtins"
	"github.com/fogo-sh/sorik/interpreter/types"
)

func Run(filename string, source []byte, args map[string]string) error {
	imagick.Initialize()
	defer imagick.Terminate()

	thread := &starlark.Thread{
		Print: func(_ *starlark.Thread, msg string) {
			log.Debug().Msg(msg)
		},
		Load: func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			return nil, fmt.Errorf("sorik does not support importing external code")
		},
	}
	thread.SetLocal("args", args)

	globals, err := starlark.ExecFile(thread, filename, source, builtins.ConstructBuiltins())
	if err != nil {
		return fmt.Errorf("error executing file: %w", err)
	}

	outputVal, found := globals["output"]
	if !found {
		return errors.New("no output value found - please ensure you assign your final image to the output variable")
	}

	outputImg, valid := outputVal.(types.Image)
	if !valid {
		return errors.New("output value must be of type Image")
	}

	err = os.WriteFile(
		fmt.Sprintf("output.%s", strings.ToLower(outputImg.Wand.GetImageFormat())),
		outputImg.Wand.GetImagesBlob(),
		0644,
	)
	if err != nil {
		return fmt.Errorf("error writing output image: %w", err)
	}

	return nil
}

func LoadAndExec(path string, args map[string]string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error loading script: %w", err)
	}

	return Run(filepath.Base(path), data, args)
}
