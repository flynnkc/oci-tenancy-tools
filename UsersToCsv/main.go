package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flynnkc/oci-tenancy-tools/UsersToCsv/pkg/makecsv"
	"github.com/flynnkc/oci-tenancy-tools/UsersToCsv/pkg/users"
)

var (
	output string
	debug  bool
)

func init() {
	flag.StringVar(&output, "out", "", "Name of file that output will be saved to")
	flag.BoolVar(&debug, "d", false, "Turn on debug mode")
	flag.Usage = func() {
		fmt.Println(flag.CommandLine.Output(), "Usage: UsersToCsv -out [filename]")
		flag.PrintDefaults()
		os.Exit(0)
	}

}

func main() {
	flag.Parse()
	makecsv.SetEnvironment(debug, output)

	if debug {
		fmt.Printf("Value of output: %v\n", output)
	}

	if output == "" && !debug {
		fmt.Println("Please specify a filename with -out")
		os.Exit(0)
	}

	// sortedUsers is a map[identityprovider][]usernames
	sortedUsers := users.GetUsers().SortByIdp()
	makecsv.MapToCsv(sortedUsers)
}
