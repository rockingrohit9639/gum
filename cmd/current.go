package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the currently active Git profile.",
	Long: `The "current" command displays the currently active Git user profile 
by reading the global Git configuration.

This includes:
  - Name 
  - Email 

Example usage:
  gum current

This helps you quickly verify which Git profile is currently applied.
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := getGitConfig("user.name")
		email, _ := getGitConfig("user.email")

		if name == "" && email == "" {
			fmt.Println("No Git profile is currently set.")
			return
		}

		fmt.Println("Currently Active Git Profile:")
		fmt.Printf("Name:  %s\n", name)
		fmt.Printf("Email: %s\n", email)
	},
}

func getGitConfig(key string) (string, error) {
	cmd := exec.Command("git", "config", "--global", key)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
