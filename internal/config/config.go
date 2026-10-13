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

const configFileName = ".gatorconfig.json"

// Helper function to get Config File Path
func getConfigFilePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDirectory, configFileName), nil
}

// Write config to JSON file
func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Make the JSON file readable
	data, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0777)

}

// SetUser method to write config structs to JSON after setting user name field
func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}
