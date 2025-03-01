package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureConfig() error {
	profileConfigPath := getProfileConfigPath()
	profileConfigDir := filepath.Dir(profileConfigPath)

	// Create directories for profiles if does not exists
	if err := os.MkdirAll(profileConfigDir, 0755); err != nil {
		return fmt.Errorf("something went wrong while creating gum directories")
	}

	// Create profiles.json if it does not exists
	if _, err := os.Stat(profileConfigPath); os.IsNotExist(err) {
		// Create the file with empty JSON object {}
		emptyJSON := []byte("{}")
		if err := os.WriteFile(profileConfigPath, emptyJSON, 0644); err != nil {
			return fmt.Errorf("something went wrong while creating profiles config file")
		}

	}

	return nil
}

/*
This function returns the absolute path to the profiles.json file
where Git user profiles are stores
Convention (~/.config/gum/profiles.json)
*/
func getProfileConfigPath() string {
	profileConfigDir := filepath.Join(os.Getenv("HOME"), ".config", "gum")
	profileConfigPath := filepath.Join(profileConfigDir, "profiles.json")

	return profileConfigPath
}
