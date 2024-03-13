package addcmd

import (
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

func isValidInput(sex string) bool {
	valid := []string{"male", "female"}
	for _, curr := range valid {
		if sex == curr {
			return true
		}
	}
	return false
}

func getAddPersonCmd() *cobra.Command {
	personCmd := &cobra.Command{
		Use:   "person",
		Short: "Add a new person",
		Long:  `This commands allows you to add a new person`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rootFlags := cmd.Flags()
			if !rootFlags.Changed("person") {
				config.PrintError("Error: Person Flag is Missing")
				os.Exit(1)
			}
			if !rootFlags.Changed("sex") {
				config.PrintError("Error: Sex Flag is Missing")
				os.Exit(1)
			}
			if !rootFlags.Changed("user") {
				config.PrintError("Error: User Flag is Missing")
				os.Exit(1)
			}
			sex := cmd.Flag("sex").Value.String()

			if !isValidInput(sex) {
				config.PrintError("Error: Wrong user sex defined")
				os.Exit(1)
			}

			user := cmd.Flag("user").Value.String()
			person := cmd.Flag("person").Value.String()

			service := repository.NewUserRepository(user)
			err := service.AddPerson(person, sex)
			if err != nil {
				config.PrintError(err)
				os.Exit(1)
			}
			config.PrintInfo("Successfully Attached Person")
		},
	}

	personCmd.Flags().String("person", "", "Add Person To User")
	personCmd.Flags().String("sex", "", "Add Sex of User")
	personCmd.Flags().String("user", "", "Associated User Flag With Person")

	return personCmd
}
