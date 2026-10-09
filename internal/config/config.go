package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Struct for .gatorconfig.json file in my home directory
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read my .gatorconfig.json file and return the Config struct
func Read() (Config, error) {

	// Retrieve users Home Directory
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	// Recreate the path to the config file
	configPath := filepath.Join(homeDirectory, ".gatorconfig.json")

	// Read the json file's contents
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	// Decode the json bytes into a Config struct
	var myConfig Config
	if err := json.Unmarshal(fileData, &myConfig); err != nil {
		return Config{}, err
	}
	return myConfig, nil
}
