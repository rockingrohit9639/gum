package cmd

import (
	"errors"
	"fmt"
	"gum/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <profile-name>",
	Short: "Add a new Git profile",
	Long:  "Add a new Git profile with a name and email storing it in the profiles.json file.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("profile name is required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		emailPrompt := promptui.Prompt{
			Label: "Email",
			Validate: func(email string) error {
				if len(email) < 3 {
					return errors.New("please enter email")
				}

				match := utils.EMAIL_REGEX.MatchString(email)
				if !match {
					return errors.New("invalid email")
				}

				return nil
			},
		}

		email, err := emailPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		namePrompt := promptui.Prompt{
			Label: "Name",
			Validate: func(name string) error {
				if len(name) < 3 {
					return errors.New("enter a valid name")
				}

				return nil
			},
		}

		name, err := namePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		newProfile := utils.Profile{
			Name:  name,
			Email: email,
		}

		utils.Profiles[profileName] = newProfile

		if err := utils.SaveProfiles(); err != nil {
			fmt.Printf("Failed to save profile %v\n", err)
			return
		}

		fmt.Printf("Successfully saved %s profile\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
