package cmd

import (
	"fmt"
	"gum/utils"
	"os/exec"

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

		profile, exists := utils.Profiles[profileName]
		if !exists {
			fmt.Printf("Error: Profile '%s' not found.\n", profileName)
			return
		}

		// set username from selected profile
		execGitConfigCmd("user.name", profile.Name)

		// set email from selected profile
		execGitConfigCmd("user.email", profile.Email)

		fmt.Printf("Switched to profile '%s' (%s)\n", profileName, profile.Email)
	},
}

func execGitConfigCmd(key, value string) {
	cmd := exec.Command("git", "config", "--global", key, value)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error setting Git config: %s -> %s\n", key, err)
	}

}

func init() {
	rootCmd.AddCommand(useCmd)
}
