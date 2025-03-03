package cmd

import (
	"fmt"
	"gum/utils"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rcCmd = &cobra.Command{
	Use:   "rc <profile>",
	Short: "rc command creates a .gumrc file in the current directory to enable automatic Git profile switching.",
	Long: `The rc command creates a .gumrc file in the current directory with the given profile name. 
This allows users to associate a Git profile with the directory.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		_, exists := utils.Profiles[profileName]
		if !exists {
			fmt.Printf("%s does not exists\n", profileName)
			return
		}

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("failed to get the current working directory")
			return
		}

		gumrcPath := filepath.Join(cwd, ".gumrc")

		// check if the .gumrc file already exists
		_, err = os.Stat(gumrcPath)
		if err == nil {
			fmt.Print(".gumrc file already exists\n")
			return
		}

		err = os.WriteFile(gumrcPath, []byte(profileName), 0644)
		if err != nil {
			fmt.Printf("could not create .gumrc file %v\n", err)
			return
		}

		fmt.Printf(".gumrc file created for %s profile\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(rcCmd)
}
