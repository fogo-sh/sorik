package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/fogo-sh/sorik/interpreter"
)

var runCmd = &cobra.Command{
	Use:   "run <script>",
	Short: "Execute a Sorik script",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := interpreter.LoadAndExec(args[0])
		if err != nil {
			log.Error().Err(err).Msg("Error executing script")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
