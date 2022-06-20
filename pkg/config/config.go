package config

import (
	"fmt"
	"os"

	"github.com/dukex/ely/pkg/cmd"
	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Path     string `yaml:"path"`
	Function string `yaml:"function"`
}

type Config struct {
	Endpoints []Endpoint `yaml:"endpoint"`
}

func Load(file string) Config {
	fmt.Println(file)
	config := Config{}

	configFileBody, err := os.ReadFile(file)
	cmd.CheckError(err)

	err = yaml.Unmarshal([]byte(configFileBody), &config)
	cmd.CheckError(err)

	return config
}
