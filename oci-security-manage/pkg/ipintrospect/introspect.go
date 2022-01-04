package introspect

import (
	"log"

	externalip "github.com/glendc/go-external-ip"
)

// GetIp takes a reference to a Logger and returns the public IP address based
// on feedback from multiple outside sources
func GetIp(log *log.Logger) string {
	// Create the default consensus,
	// using the default configuration and no logger.
	consensus := externalip.DefaultConsensus(nil, log)

	// By default Ipv4 or Ipv6 is returned,
	// use the function below to limit yourself to IPv4,
	// or pass in `6` instead to limit yourself to IPv6.
	consensus.UseIPProtocol(4)

	// Get your IP,
	// which is never <nil> when err is <nil>.
	ip, err := consensus.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	return ip.String()
}
