package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PostgresConnString string `yaml:",omitempty" mapstructure:"postgresConnString"`
	MigrationsFolder   string `yaml:",omitempty" mapstructure:"migrationsFolder"`
	Port               int    `yaml:",omitempty" mapstructure:"port"`
}

func FromFile(file string) (Config, error) {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("read file: %w", err)
	}

	var m map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return Config{}, fmt.Errorf("yaml unmarshal : %w", err)
	}

	var c Config
	err = mapstructure.Decode(m, &c)
	if err != nil {
		return Config{}, fmt.Errorf("yaml unmarshal : %w", err)
	}

	return c, nil
}
