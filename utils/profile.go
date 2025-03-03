package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Profiles = map[string]Profile{}

/*
This function loads the profiles in &Profiles
*/
func LoadProfiles() error {
	profilesConfigPath := getProfileConfigPath()

	data, err := os.ReadFile(profilesConfigPath)
	if err != nil {
		return err
	}

	// Parse and load profiles.json in &Profiles
	if err := json.Unmarshal(data, &Profiles); err != nil {
		return err
	}

	return nil
}

/*
This function saves the &Profiles in profiles.json
*/
func SaveProfiles() error {
	profilesConfigPath := getProfileConfigPath()

	data, err := json.MarshalIndent(Profiles, "", "    ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(profilesConfigPath, data, 0644); err != nil {
		return err
	}

	return nil
}

/*
This function applies the given profile
*/
func ApplyProfile(profileName string) error {
	profile, exists := Profiles[profileName]
	if !exists {
		return fmt.Errorf("profiles %s not found", profileName)
	}

	// set username from selected profile
	execGitConfigCmd("user.name", profile.Name)

	// set email from selected profile
	execGitConfigCmd("user.email", profile.Email)

	return nil
}

func execGitConfigCmd(key, value string) {
	cmd := exec.Command("git", "config", "--global", key, value)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error setting Git config: %s -> %s\n", key, err)
		return
	}
}
