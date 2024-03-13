package connectcmd

import (
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

var connectCmd *cobra.Command

func GetConnectCmd() *cobra.Command {
	return connectCmd
}

func init() {
	connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Use this to define relationships",
		Long:  "Use this command to define the type relationship between two person",
		Args:  cobra.ExactArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			rootFlags := cmd.Flags()
			if !rootFlags.Changed("user") {
				config.PrintError("Error: User Flag is Missing")
				os.Exit(1)
			}
			user := cmd.Flag("user").Value.String()

			member := args[0]
			relationship := args[2]
			relative := args[4]

			defineRelationship(user, member, relationship, relative)
		},
	}

	connectCmd.Flags().String("user", "", "User defined relationship")

}

func defineRelationship(currentUser string, memberName string, relationship string, dependent string) {
	r := repository.NewUserRepository(currentUser)
	err := r.AttachRelationship(memberName, relationship, dependent)

	if err != nil {
		config.PrintError(err)
		os.Exit(1)
	}
	config.PrintInfo("Successfully Attached relationship")
}
