package handleyaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration contains all data in the config.yaml file
type Configuration struct {
	Version        int    `yaml:"version,omitempty"`
	ConfigLocation string `yaml:"-"`
	OciDirectory   string `yaml:"ociDirectory"`
	LastIp         string `yaml:"lastIp,omitempty"`
	Resources      []struct {
		Name    string `yaml:"name"`
		Profile string `yaml:"profile"`
		Type    string `yaml:"type"`
		OCID    string `yaml:"ocid"`
		Id      string `yaml:"id,omitempty"`
	}
}

// NewConfig initializes an empty Configuration struct and returns a pointer
func NewConfig() *Configuration {
	c := &Configuration{}
	return c
}

// DefaultConfig returns a loaded Configuration struct with default config
// file location
func DefaultConfig() *Configuration {
	c := NewConfig()
	c.ConfigLocation = "./config.yaml"

	err := c.ReadConfigFile("./config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
	return c
}

// ReadConfigFile attempts to load a config file into a Configuration struct
func (c *Configuration) ReadConfigFile(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

// WriteConfig writes overwrites the config.yaml file with updated yaml
func (c *Configuration) WriteConfig() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("Error writing to config.yaml file: %v\n", err)
	}
	log.Println(string(data))
}
