package main

import (
	"flag"
	"log"
	"os"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/cloud"
	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
	introspect "github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/ipintrospect"
)

var debug bool
var logger *log.Logger

func init() {
	flag.BoolVar(&debug, "d", false, "Enter debug mode for troubleshooting")
}

func main() {
	flag.Parse()

	logger = log.Default()
	setEnvironments()

	// Get Configuration struct containing information in config.yaml
	c := handleyaml.DefaultConfig()

	// If IP matches previous IP, do nothing and exit
	ip := introspect.GetIp(logger)
	if debug {
		logger.Printf("[DEBUG] External IP returned: %v\n[DEBUG] Last recorded IP: %v\n",
			ip, c.LastIp)
	}
	if c.LastIp == ip && !debug {
		logger.Println("Nothing to do here")
		os.Exit(0)
	}

	cloud.SetConfig(c)
	cloud.ReadConfig()

	// Set LastIp to returned IP before exiting
	logger.Printf("INFO: Updating to match current IP %v\n", ip)
	c.LastIp = ip
	if err := c.WriteConfig(); err != nil {
		logger.Fatalf("ERROR: Failed to write updates to config.yaml: %v", err)
	}
	os.Exit(0)
}

func setEnvironments() {
	cloud.SetEnvironment(debug, logger)
	handleyaml.SetEnvironment(debug, logger)
}
