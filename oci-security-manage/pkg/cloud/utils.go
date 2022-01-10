package cloud

import "github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"

// convertConfigToResource consumes a Configruation and returns a slice of
// resources
func convertConfigToResource(c *handleyaml.Configuration) []resource {
	resources := make([]resource)
	for _, r := range c.Resources {
		s := resource{
			name:    r.Name,
			profile: r.Profile,
			object:  r.Type,
			ocid:    r.OCID,
			id:      r.Id,
			region:  r.Region,
		}
		resources = append(resources, s)
	}
	return resources
}

// sortResources sorts a slice of resources and returns a map of slices of
// resources sorted by region
func sortResources(r []resource) map[string][]resource {
	m := make(map[string][]resource)
	for _, s := range r {
		m[s.region] = append(m[s.region], s)
	}
	return m
}
