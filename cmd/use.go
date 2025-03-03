package cmd

import (
	"fmt"
	"gum/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch to a saved Git profile",
	Long: `The 'use' command allows you to switch between different Git profiles 
saved in your configuration file. It updates the global Git user settings 
(user.name, user.email) based on the selected profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		existingProfileNames := []string{}
		for profileName := range utils.Profiles {
			existingProfileNames = append(existingProfileNames, profileName)
		}

		profileNamePrompt := promptui.Select{
			Label: "Select profile",
			Items: existingProfileNames,
		}

		_, profileName, err := profileNamePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = utils.ApplyProfile(profileName)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("Switched to profile '%s'\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
