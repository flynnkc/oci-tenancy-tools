package handleyaml

import (
	"log"

	"gopkg.in/yaml.v3"
)

// WriteConfig writes to a config file
func (c *Configuration) WriteConfig() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
