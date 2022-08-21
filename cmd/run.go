package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.starlark.net/starlark"

	"github.com/fogo-sh/sorik/interpreter"
)

var runCmd = &cobra.Command{
	Use:   "run <script>",
	Short: "Execute a Sorik script",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scriptArgs, err := cmd.Flags().GetStringToString("args")
		if err != nil {
			log.Error().Err(err).Msg("Error getting script args")
			os.Exit(1)
		}

		err = interpreter.LoadAndExec(args[0], scriptArgs)
		if err != nil {
			evalErr, isEvalErr := errors.Unwrap(err).(*starlark.EvalError)
			if isEvalErr {
				fmt.Println(evalErr.Backtrace())
			} else {
				log.Error().Err(err).Msg("Error executing script")
			}

			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringToString("args", map[string]string{}, "arguments to make available to the executed Sorik script")
}
