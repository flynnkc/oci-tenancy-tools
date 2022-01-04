package handleyaml

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const comment string = `# This config.yaml file contains information required to update selected OCI 
# Security List and Network Security Group resources.

`

// Configuration contains all data in the config.yaml file
type Configuration struct {
	Version        int    `yaml:"version"`
	configLocation string `yaml:"-"`
	OciDirectory   string `yaml:"ociDirectory"`
	LastIp         string `yaml:"lastIp"`
	Resources      []struct {
		Name    string `yaml:"name"`
		Profile string `yaml:"profile"`
		Type    string `yaml:"type"`
		OCID    string `yaml:"ocid"`
		Id      string `yaml:"id"`
		Region  string `yaml:"region"`
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

	// Set default file path to .oci/config.yaml or ./config.yaml if no ~/.oci
	// directory exists
	file := setDefaultFileLocation()

	err := c.ReadConfigFile(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
	return c
}

// ReadConfigFile attempts to load a config file into a Configuration struct
func (c *Configuration) ReadConfigFile(file string) error {
	var err error
	// Set config file location
	c.configLocation, err = filepath.Abs(file)
	if err != nil {
		return err
	}

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
func (c *Configuration) WriteConfig() error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	// Append comments to data before writing
	data = append([]byte(comment), data...)

	os.WriteFile(c.configLocation, data, 0640)
	return nil
}