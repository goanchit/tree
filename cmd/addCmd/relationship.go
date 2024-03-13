package addcmd

import (
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

func getAddRelationshipCmd() *cobra.Command {
	relationshipCmd := &cobra.Command{
		Use:   "relationship",
		Short: "Add a new relationship",
		Long:  `This commands allows you to add a new relationship`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rootFlags := cmd.Flags()
			if !rootFlags.Changed("relationship") {
				config.PrintError("Error: Person Flag is Missing")
				os.Exit(1)
			}
			if !rootFlags.Changed("user") {
				config.PrintError("Error: User Flag is Missing")
				os.Exit(1)
			}

			user := cmd.Flag("user").Value.String()
			relation := cmd.Flag("relationship").Value.String()

			service := repository.NewUserRepository(user)
			err := service.AddRelation(relation)
			if err != nil {
				config.PrintError(err)
				os.Exit(1)
			}
			config.PrintInfo("Successfully Attached Relationship")
		},
	}

	relationshipCmd.Flags().String("relationship", "", "Define User Relationships")
	relationshipCmd.Flags().String("user", "", "Associated User Flag With Relationship")

	return relationshipCmd
}
