package queriescmd

import (
	"family-tree/config"
	"family-tree/repository"
	"os"

	"github.com/spf13/cobra"
)

var queriesCmd *cobra.Command

func GetQueriesCmd() *cobra.Command {
	return queriesCmd
}

func init() {

	queriesCmd = &cobra.Command{
		Use:   "relations",
		Short: "Relation To Member",
		Long:  "Gets Relation To Member",
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
			records, err := r.FindRecords(relation, name, true)

			var memberIds []uint
			if err != nil {
				config.PrintError(err)
				os.Exit(1)
			}
			if len(records) == 0 {
				config.PrintInfo("No Relation to the member defined. Finding reverse relation")
			} else {
				for _, val := range records {
					id := val.RelativeID
					memberIds = append(memberIds, id)
				}
			}

			reverseRelationshipMap := map[string][]string{
				"father":   {"son", "daughter"},
				"mother":   {"son", "daughter"},
				"son":      {"father", "mother"},
				"daughter": {"father", "mother"},
			}

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

			if len(memberIds) == 0 {
				config.PrintInfo("No Match Found!!!")
			} else {
				config.PrintInfo("memberIds found", memberIds)
				res, _ := r.FindById(memberIds)
				for _, val := range res {
					config.PrintInfo("relations found: ", val.Name)
				}
			}

		},
	}
	queriesCmd.Flags().String("user", "", "Associated User Flag With Person")
}
