package builtins

import (
	"fmt"

	"go.starlark.net/starlark"
)

func getArg(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		name string
	)

	if err := starlark.UnpackArgs(
		fn.Name(), args, kwargs,
		"name", &name,
	); err != nil {
		return nil, err
	}

	scriptArgs := thread.Local("args").(map[string]string)

	val, found := scriptArgs[name]
	if !found {
		return nil, fmt.Errorf("unknown script argument %s", name)
	}

	return starlark.String(val), nil
}
