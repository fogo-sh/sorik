package interpreter

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"

	"github.com/fogo-sh/sorik/interpreter/builtins"
	"github.com/fogo-sh/sorik/interpreter/types"
)

func Run(filename string, source []byte, args map[string]string) ([]byte, error) {
	resolve.AllowGlobalReassign = true

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
		return nil, fmt.Errorf("error executing file: %w", err)
	}

	outputVal, found := globals["output"]
	if !found {
		return nil, errors.New("no output value found - please ensure you assign your final image to the output variable")
	}

	outputImg, valid := outputVal.(types.Image)
	if !valid {
		return nil, errors.New("output value must be of type Image")
	}

	return outputImg.Wand.GetImagesBlob(), nil
}

func LoadAndExec(path string, args map[string]string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error loading script: %w", err)
	}

	return Run(filepath.Base(path), data, args)
}
