package countcmd

import (
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

var countCmd *cobra.Command

func GetCountCmd() *cobra.Command {
	return countCmd
}

func init() {
	countCmd = &cobra.Command{
		Use:   "count",
		Short: "Get Count Of Relations",
		Long:  "Get Count of Relations",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			rootFlags := cmd.Flags()
			if !rootFlags.Changed("user") {
				config.PrintError("Error: Person Flag is Missing")
				os.Exit(1)
			}
			user := cmd.Flag("user").Value.String()

			relation := args[0]
			name := args[2]

			r := repository.NewUserRepository(user)
			records, err := r.FindRecords(relation, name, false)
			if err != nil {
				config.PrintError(err)
				os.Exit(1)
			}
			config.PrintInfo(len(records), ": records found")

			reverseRelationshipMap := map[string][]string{
				"father":   {"son", "daughter"},
				"mother":   {"son", "daughter"},
				"son":      {"father", "mother"},
				"daughter": {"father", "mother"},
			}

			var memberIds []uint

			relationshipKeys := reverseRelationshipMap[relation]
			for _, val := range relationshipKeys {
				records, err := r.FindRecords(val, name, false)
				if err != nil {
					config.PrintError(err)
					os.Exit(1)
				}
				if len(records) > 0 {
					for _, val := range records {
						id := val.RelativeID
						memberIds = append(memberIds, id)
					}
				}
			}
			config.PrintInfo(len(memberIds), ": records found")
		},
	}
	countCmd.Flags().String("user", "", "Associated User Flag With Person")
}
