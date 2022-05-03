package cloud

import (
	"log"
	"sync"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
)

var debug = false
var logger *log.Logger

// SetEnvironment sets the environment variables shared by the application
func SetEnvironment(d bool, l *log.Logger) {
	debug, logger = d, l
}

// PrintConfig prints config TODO parse data to make useful e.g. using creds x to update y
func PrintConfig(config *handleyaml.Configuration) {
	logger.Printf("[INFO] Config file data in cloud package: %v\n", &config)
}

// UpdateResources is control function for resource updates
func UpdateResources(c *handleyaml.Configuration, ip string) error {
	rsources := convertConfigToResource(c, ip)
	if debug {
		logger.Printf("[DEBUG]: Config unpacked into resources: %v\n", rsources)
	}

	c = c.NewConfigFromBase()

	// Loop through rsources to update
	var wg sync.WaitGroup
	for i, r := range rsources {
		if debug {
			logger.Printf("[DEBUG]: Resource %d: %v\n", i, r)
		}
		wg.Add(1)

		// Fire off goroutines to update resources using anonymous function
		go func(obj *resource) {
			defer wg.Done()
			switch obj.object {
			case "NSG":
				logger.Printf("[INFO] Updating NSG: %v\n", obj.name)
				err := updateNsg(obj)
				if err != nil {
					logger.Printf("[ERROR]: Failed to update %v: %v", obj.name, err)
				}
			default:
				logger.Printf("[WARN] Unsupported resource type %v on resource %v",
					obj.object, obj.name)
				return
			}
			// Append resource to new base config object
			// Is this a good idea? I already have a reference to c
			appendResourceToConfig(c, obj)
		}(&r)
	}

	wg.Wait()

	return nil
}
