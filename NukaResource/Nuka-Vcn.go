package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/v45/common"
	"github.com/oracle/oci-go-sdk/v45/core"
)

func main() {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	checkFatalErr(err)

	req := core.GetVcnTopologyRequest{
		VcnId:         common.String("ocid1.vcn.oc1.iad.amaaaaaac3adhhqaxlxp33czi6u3u3z2cjk44xrxouslp5bdvckntxz2hqoq"),
		CompartmentId: common.String("ocid1.compartment.oc1..aaaaaaaa2dysu5uvi6by4ie5hzzt5cyn6kzjz3s2oawnlyoptprbca2k6mla"),
	}

	resp, err := client.GetVcnTopology(context.Background(), req)
	checkFatalErr(err)

	fmt.Println(resp)
}

func checkFatalErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
