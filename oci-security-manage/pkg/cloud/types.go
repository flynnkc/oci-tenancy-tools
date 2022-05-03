package cloud

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/core"
)

type resource struct {
	name     string
	ip       string
	profile  string
	object   string // Security List or NSG, can't use type as it's a keyword
	ocid     string
	id       string
	region   string
	port     int
	protocol string
	lastIp   string
}

// setResourceIp does exactly what you think it does
func (r *resource) setResourceIp(ip string) {
	r.ip = ip
}

// setResourceid sets a resource ID
func (r *resource) setResourceId(id string) {
	r.id = id
}

// getNsgId gets and sets the ID for the NSG Rule
func (r *resource) getNsgSecurityRuleId(client core.VirtualNetworkClient) error {
	// Create request struct
	req := core.ListNetworkSecurityGroupSecurityRulesRequest{
		NetworkSecurityGroupId: common.String(r.ocid),
		Direction:              "INGRESS",
		Limit:                  common.Int(200),
	}

	// Create context for request with 3 second timeout
	ctx, cancel := setContextTimeoutSeconds(3)
	defer cancel()

	// Send request and get response
	resp, err := client.ListNetworkSecurityGroupSecurityRules(ctx, req)
	if err != nil {
		return fmt.Errorf("[ERROR]: Error during list security rules operation: %v\n",
			err)
	} else {
		if debug {
			logger.Printf("[DEBUG]: Security rule ID query result: %v\n", resp)
			logger.Printf("[DEBUG]: Rule values from config: Last IP:%v -- Port: %v\n", r.lastIp, r.port)
		}
		for _, item := range resp.Items {
			if *item.Source == (r.lastIp+"/32") && *item.TcpOptions.SourcePortRange.Max == r.port {
				setResourceId(*item.Id)
			}
		}
		if r.id == "" {
			return fmt.Errorf("Could not find ID for %v\n", r.name)
		}
		return nil
	}
}
