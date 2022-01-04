package main

import (
	"fmt"

	"github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/handleyaml"
	//	introspect "github.com/flynnkc/oci-tenancy-tools/oci-security-manage/pkg/ipintrospect"
)

func main() {
	c := handleyaml.DefaultConfig()
	fmt.Println(&c, "\n")

	c.WriteConfig()
}
