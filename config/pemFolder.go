package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func GetPemFolder() (string, error) {
	pemFolderPath := viper.GetString(PemFolderKey)
	if pemFolderPath != "" {
		return pemFolderPath, nil
	} else {
		return pemFolderPath, errors.New("config file: " + PemFolderKey + " key is missing")
	}
}

func SetConfigPemFolder(value string) error {
	_, err := os.ReadDir(value)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(value)
	if err != nil {
		return err
	}

	yamlContent, err := os.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return err
	}
	configContent := ConfigFile{}

	err = yaml.Unmarshal(yamlContent, &configContent)
	if err != nil {
		return err
	}

	configContent.PemFolder = absPath

	yamlContent, err = yaml.Marshal(configContent)
	if err != nil {
		return err
	}

	err = os.WriteFile(viper.ConfigFileUsed(), yamlContent, 0640)
	if err != nil {
		return err
	}

	return nil
}
