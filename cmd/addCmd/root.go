package addcmd

import (
	"github.com/spf13/cobra"
)

func GetAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Use this to add members/relationships",
		Long:  "This command is used to add members/relationships",
	}

	addCmd.AddCommand(getAddPersonCmd())
	addCmd.AddCommand(getAddRelationshipCmd())
	return addCmd
}
