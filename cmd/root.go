package cmd

import (
	queriescmd "family-tree/cmd/queriesCmd"
	usercmd "family-tree/cmd/userCmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "Family tree generator",
	Long:  "A family tree generator is used to generate family tree. User can add members, add relations",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(
		usercmd.UserCmd,
		queriescmd.GetQueriesCmd(),
	)
}
