package utils

import (
	"encoding/json"
	"os"
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
