package handleyaml

// util.go is for functions internal to the handleyaml package

import (
	"os"
	"os/user"
	"path/filepath"
)

func setDefaultFileLocation() string {
	usr, err := user.Current()
	if err != nil {
		return "./config"
	}

	hd := filepath.Join(usr.HomeDir, ".oci/config.yaml")

	// Test if .oci directory exists
	_, err = os.Stat(hd)
	if os.IsNotExist(err) {
		return "./config"
	}

	return hd
}
