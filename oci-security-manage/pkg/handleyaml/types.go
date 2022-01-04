package handleyaml

// SecurityObject is the yaml data needed to update Security List and Network
// Security Group objects
type SecurityObject struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile,omitempty"`
	Type    string `yaml:"type"`
	OCID    string `yaml:"ocid"`
	Id      string `yaml:"id"`
}

// Configuration contains all data in the config.yaml file
type Configuration struct {
	Version   int    `yaml:"version,omitempty"`
	Directory string `yaml:"directory"`
	LastIp    string `yaml:"lastip,omitempty"`
	Resources []struct {
		Name    string `yaml:"name"`
		Profile string `yaml:"profile"`
		Type    string `yaml:"type"`
		OCID    string `yaml:"ocid"`
		Id      string `yaml:"id,omitempty"`
	}
}
