package interpreter

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func Run(filename string, source []byte) error {
	imagick.Initialize()
	defer imagick.Terminate()

	thread := &starlark.Thread{
		Print: func(_ *starlark.Thread, msg string) {
			log.Debug().Msg(msg)
		},
	}

	globals, err := starlark.ExecFile(thread, filename, source, constructBuiltins())
	if err != nil {
		return fmt.Errorf("error executing file: %w", err)
	}

	spew.Dump(globals)

	return nil
}

func LoadAndExec(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error loading script: %w", err)
	}

	return Run(filepath.Base(path), data)
}
