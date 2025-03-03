package cmd

import (
	"fmt"
	"gum/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved Git profiles.",
	Long: `The 'list' command displays all saved Git profiles from the configuration file. 
Each profile includes details such as the profile name, associated email. 
This helps users quickly view available profiles and switch between them as needed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(utils.Profiles) < 1 {
			fmt.Println("There is not profile yet.")
			return
		}

		for profileName, profile := range utils.Profiles {
			fmt.Println("----------------------")
			fmt.Printf("Profile: %s\n", profileName)
			fmt.Printf("Email: %s\n", profile.Email)
			fmt.Printf("Name: %s\n", profile.Name)
			fmt.Println("----------------------")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
