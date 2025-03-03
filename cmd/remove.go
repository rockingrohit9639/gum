package cmd

import (
	"fmt"
	"gum/utils"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <profile>",
	Short: "Remove a saved Git profile from gum.",
	Long: `The "remove" command allows you to delete a saved Git profile from gum. 
Once removed, the profile will no longer be available for switching.

Example usage:
  gum remove my-profile

This action is irreversible, so use it with caution.
You may be prompted for confirmation before deletion.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		_, exists := utils.Profiles[profileName]
		if !exists {
			fmt.Printf("%s does not exists\n", profileName)
			return
		}

		delete(utils.Profiles, profileName)
		utils.SaveProfiles()

		fmt.Printf("%s removed successfully\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
