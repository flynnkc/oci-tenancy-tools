package cloud

import (
	"context"
	"fmt"
	"time"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/core"
)

// convertConfigToResource consumes a Configruation and returns a slice of
// resources
func convertConfigToResource(c *handleyaml.Configuration, ip string) []resource {
	resources := make([]resource, 0)
	for _, r := range c.Resources {
		s := resource{
			name:     r.Name,
			ip:       ip,
			profile:  r.Profile,
			object:   r.Type,
			ocid:     r.OCID,
			id:       r.Id,
			region:   r.Region,
			port:     r.Port,
			protocol: r.Protocol,
			lastIp:   c.LastIp,
		}
		resources = append(resources, s)
	}
	return resources
}

// setContextTimeoutSeconds creates and returns a context with specified
// in seconds
func setContextTimeoutSeconds(s int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s)*time.Second)
	return ctx, cancel
}

// sortResources sorts a slice of resources and returns a map of slices of
// resources sorted by region
// May not be required
func sortResources(r []resource) map[string][]resource {
	m := make(map[string][]resource)
	for _, s := range r {
		m[s.region] = append(m[s.region], s)
	}
	return m
}

// updateNsg takes an NSG object and performs an update operation
func updateNsg(nsg resource) error {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		return err
	}

	// Get NSG ID if it is not present
	if nsg.id == "" {
		err = nsg.getNsgSecurityRuleId(client)
		if err != nil {
			return err
		}
	}

	nsgRuleDetails := core.UpdateNetworkSecurityGroupSecurityRulesDetails{
		SecurityRules: []core.UpdateSecurityRuleDetails{
			core.UpdateSecurityRuleDetails{
				Direction:  core.UpdateSecurityRuleDetailsDirectionIngress,
				Id:         common.String(nsg.id),
				Protocol:   common.String(nsg.protocol),
				Source:     common.String(nsg.ip + "/32"),
				SourceType: core.UpdateSecurityRuleDetailsSourceTypeCidrBlock,
			},
		},
	}

	switch nsg.protocol {
	case "6":
		nsgRuleDetails.SecurityRules[0].TcpOptions = &core.TcpOptions{SourcePortRange: &core.PortRange{Max: common.Int(nsg.port), Min: common.Int(nsg.port)}}
	case "17":
		nsgRuleDetails.SecurityRules[0].UdpOptions = &core.UdpOptions{SourcePortRange: &core.PortRange{Max: common.Int(nsg.port), Min: common.Int(nsg.port)}}
	default:
		return fmt.Errorf("[ERROR]: Protocol value %v not recognized, port %v not updated\n",
			nsg.protocol, nsg.name)
	}

	req := core.UpdateNetworkSecurityGroupSecurityRulesRequest{
		NetworkSecurityGroupId:                         common.String(nsg.ocid),
		UpdateNetworkSecurityGroupSecurityRulesDetails: nsgRuleDetails,
	}

	_, err = client.UpdateNetworkSecurityGroupSecurityRules(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
