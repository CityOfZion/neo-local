package config

import (
	"fmt"
	"log"
	"os"

	"github.com/atrox/homedir"
	"github.com/spf13/viper"
)

const (
	configDir = "~/.neo-local"
)

// NewConfig creates a new set of config variables to be used by neo-local
func NewConfig(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	viper.SetConfigFile(filename)

	log.Printf(" %s", filename)

	err := viper.ReadInConfig()
	if err != nil {
		viper.Set("initialSetup", true)
		viper.SafeWriteConfigAs(filename)
	}

	return nil
}

// DirPath retrieves the config folder complete path
func DirPath() (string, error) {
	directory, err := homedir.Expand(configDir)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			return "", err
		}
	}

	return directory, nil
}

// File retrieves the config file complete path
func File() (string, error) {
	dirPath, err := DirPath()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/config.json", dirPath), nil
}
