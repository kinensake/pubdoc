package command

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "pubdoc",
	Short: "A tool to create docusaurus from epub",
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(addCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
