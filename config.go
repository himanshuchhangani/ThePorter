package theporter

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds all the configuration required to up and
// run theporter
type Config struct {
	DBConfig *DBConfig `yaml:"Database"`
}

// DBConfig holds all the configurations of DB for authservice server.
type DBConfig struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

// ReadConfig reads the YAML config and uses it to initialize the Config for theporter.
func ReadConfig(confFile string) (*Config, error) {

	confReader, err := os.Open(confFile)
	if err != nil {
		return nil, fmt.Errorf("missing config file %s error : %v", confFile, err)
	}
	defer confReader.Close()

	var config Config
	yamlData := yaml.NewDecoder(confReader)
	yamlData.SetStrict(true)
	err = yamlData.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	return &config, nil
}
