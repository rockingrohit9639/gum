package cmd

import (
	"fmt"
	"gum/utils"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Activates the Git profile specified in the .gumrc file of the current directory.",
	Long: `The auto command reads the .gumrc file in the current directory and switches to the Git profile defined in it. This allows users to quickly apply a pre-set profile for the project without manually switching.

Usage:

gum auto`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Print("could not get current directory\n")
			return
		}

		gumrcPath := filepath.Join(cwd, ".gumrc")

		_, err = os.Stat(gumrcPath)
		if os.IsNotExist(err) {
			fmt.Print(".gumrc file does not exists in current directory\n")
			return
		}

		data, err := os.ReadFile(gumrcPath)
		if err != nil {
			fmt.Print("could not read .gumrc file")
			return
		}

		profileName := string(data)

		err = utils.ApplyProfile(profileName)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("%s applied successfully\n", profileName)
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)
}
