package command

import (
	"github.com/kinensake/pubdoc/internal/docusaurus"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:                   "new [project-name]",
	Short:                 "Create a new docusaurus project",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run:                   newCmdHandler,
}

func newCmdHandler(cmd *cobra.Command, args []string) {
	err := docusaurus.NewProject(args[0])
	if err != nil {
		cobra.CompErrorln(err.Error())
	}
}

var addCmd = &cobra.Command{
	Use:                   "add [epub-path]",
	Short:                 "Add an epub to the project",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run:                   addCmdHandler,
}

func addCmdHandler(cmd *cobra.Command, args []string) {
	err := docusaurus.AddEpub(args[0])
	if err != nil {
		cobra.CompErrorln(err.Error())
	}
}
