package usercmd

import (
	addcmd "family-tree/cmd/addCmd"
	connectcmd "family-tree/cmd/connectCmd"
	countcmd "family-tree/cmd/countCmd"
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

var UserCmd = &cobra.Command{
	Use:   "login",
	Short: "Use this command to login as user. Pass with a flag",
	Long:  "This command is used to impersonate to a existing user or creates a new one. This is useful to save the record",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		rootFlags := cmd.Flags()
		if !rootFlags.Changed("user") {
			config.PrintError("Error: User Flag is Missing")
			os.Exit(1)
		}

		user := cmd.Flag("user").Value.String()
		loginUser(user)
	},
}

func init() {
	UserCmd.AddCommand(
		addcmd.GetAddCmd(),
		connectcmd.GetConnectCmd(),
		countcmd.GetCountCmd(),
	)
	UserCmd.PersistentFlags().String("user", "", "Current User")
}

func loginUser(userName string) {
	repo := repository.NewUserRepository(userName)
	if err := repo.LoginIntoUser(); err != nil {
		config.PrintError("failed to login into user")
		os.Exit(1)
	}
	config.PrintInfo("Successfully logged in")
}
