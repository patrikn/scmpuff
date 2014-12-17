package inits

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Since the flags are defined and used in different locations, we need to
// define a variable outside with the correct scope to assign the flag to work
// with.
var includeAliases bool
var outputScript bool

// CommandInit generates the command handler for `scmpuff init`
func CommandInit() *cobra.Command {

	var InitCmd = &cobra.Command{
		Use:   "init",
		Short: "Output initialization script",
		Long: `
Output the bash/zsh initialization script for scmpuff.

This should probably be evaluated in your shell startup.
    `,
		Run: func(cmd *cobra.Command, args []string) {
			if outputScript {
				printScript()
			} else {
				fmt.Println(helpString())
			}
		},
	}

	// --aliases
	InitCmd.Flags().BoolVarP(
		&includeAliases,
		"aliases", "a", true,
		"Include short aliases for convenience",
	)

	// --show
	InitCmd.Flags().BoolVarP(
		&outputScript,
		"show", "s", false,
		"Output scmpuff initialization scripts",
	)

	return InitCmd
}
