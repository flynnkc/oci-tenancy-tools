package cloud

import (
	"log"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
)

var config *handleyaml.Configuration
var debug = false
var logger *log.Logger

// SetEnvironment sets the environment variables shared by the application
func SetEnvironment(d bool, l *log.Logger) {
	debug, logger = d, l
}

// SetConfig sets the Configuration struct in the cloud package
func SetConfig(c *handleyaml.Configuration) {
	config = c
}

// ReadConfig prints config
func ReadConfig() {
	logger.Printf("[INFO] Config file data in cloud package: %v\n", *config)
}
