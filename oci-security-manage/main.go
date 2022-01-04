package main

import (
	"log"
	"os"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
	introspect "github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/ipintrospect"
)

func main() {
	logger := log.Default()
	c := handleyaml.DefaultConfig()

	// If IP matches previous IP, do nothing and exit
	ip := introspect.GetIp(logger)
	if c.LastIp == ip {
		logger.Println("Nothing to do here")
		os.Exit(0)
	}

	logger.Printf("INFO: Updating to match current IP %v\n", ip)
	c.LastIp = ip
	if err := c.WriteConfig(); err != nil {
		logger.Fatalf("ERROR: Failed to write updates to config.yaml: %v", err)
	}
	os.Exit(0)
}
