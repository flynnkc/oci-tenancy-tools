package users

import (
	"log"

	"github.com/oracle/oci-go-sdk/v36/common"
	"github.com/oracle/oci-go-sdk/v36/identity"
)

var cp common.ConfigurationProvider

func setConfigProvider(profile string) {
	switch profile {
	case "DEFAULT":
		cp = common.DefaultConfigProvider()
	default:
		log.Fatal("No profile specified in users.setConfigProvider")
	}
}

// makeConfigurationProvider is used to generate a config file to authenticate with OCI
// using default file at $HOME/.oci/config
func makeDefaultConfigurationProvider() (identity.IdentityClient, error) {
	setConfigProvider("DEFAULT")
	c, err := identity.NewIdentityClientWithConfigurationProvider(cp)

	return c, err
}
