package config

import (
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	PemFolder string `yaml:"PEM_FOLDER"`
}

// create the default .remoteSSH.yaml config file and pems/ directory
func createNewConfigFile(home string) {
	configFullPath := home + "/" + ".remoteSSH.yaml"

	_, err := os.Create(configFullPath)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	configExample := ConfigFile{PemFolder: home + "/pems"}

	yamlExample, err := yaml.Marshal(&configExample)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = os.WriteFile(configFullPath, yamlExample, 0640)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = os.Mkdir(home+"/pems", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	configFullPath := home + "/" + ".remoteSSH.yaml"

	_, err = os.ReadFile(configFullPath)
	if err != nil {
		// If config file does not exist, create one
		createNewConfigFile(home)
	}

	// Search config in home directory with name ".remoteSSH_config".
	viper.AddConfigPath(home)
	viper.SetConfigName(".remoteSSH.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func Main() {
	initConfig()

	_, err := GetPemFolder()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
