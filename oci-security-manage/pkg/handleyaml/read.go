package handleyaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// GetConfig reads, unmarshals, and returns a reference to a Configuration type
func GetConfig() *Configuration {
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Fatalf("Error reading data from config file: %v\n", err)
	}

	var c Configuration

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("Error unmarshaling yaml: %v\n", err)
	}

	return &c
}
