package cmd

import (
	"fmt"
	"gum/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gum",
	Short: "A CLI tool to manage multiple Git profiles and configurations.",
	Long: `Gum (Git User Manager) is a lightweight command-line tool that helps you manage 
multiple Git profiles with ease. Whether you work with multiple repositories, 
different accounts, or need to switch Git configurations frequently, Gum makes it simple.

Features:
- Add and manage multiple Git profiles
- Instantly switch between profiles with a single command
- Store default Git configurations and apply them globally
- View and modify your current Git identity
- Works with JSON-based config files for easy manual edits

Usage:
  gum [command]

Available Commands:
  add         Add a new Git profile
  list        List all saved profiles
  use         Switch to a specific profile
  remove      Remove a profile
  current     Show the active Git profile

Use "gum [command] --help" for more information about a command.
`,
}

func init() {
	if err := utils.EnsureConfig(); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
