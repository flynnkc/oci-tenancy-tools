package users

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v36/common"
	"github.com/oracle/oci-go-sdk/v36/identity"
)

// Users is a slice of identity.User
type Users []identity.User

// GetUsers retrieves the list of users from all identity providers in the tenancy
func GetUsers() Users {
	var u Users
	client, err := makeDefaultConfigurationProvider()
	if err != nil {
		log.Fatal(err)
	}

	tenant, err := cp.TenancyOCID() // Get tenancy OCID for use in request struct
	if err != nil {
		log.Fatal(err)
	}

	request := identity.ListUsersRequest{
		CompartmentId:  common.String(tenant),
		LifecycleState: identity.UserLifecycleStateActive, // Only active users
	}

	// Closure to pass into for loop
	getUsersFunc := func(id identity.ListUsersRequest) (identity.ListUsersResponse, error) {
		return client.ListUsers(context.Background(), id)
	}

	for response, err := getUsersFunc(request); ; response, err = getUsersFunc(request) {
		if err != nil {
			log.Fatal(err)
		}

		// Be a good internet denizen and rate limit
		time.Sleep(time.Second)

		u = append(u, response.Items...)

		if response.OpcNextPage != nil {
			request.Page = response.OpcNextPage
		} else {
			break
		}
	}

	return u
}

// SortByIdp sorts users by identity domain and returns map of strings
func (u Users) SortByIdp() map[string][]string {
	sorted := make(map[string][]string)

	for _, user := range u {
		name := *user.Name
		split := strings.Split(name, "/")
		if len(split) == 1 {
			sorted["local"] = append(sorted["local"], split[0])
		} else {
			sorted[split[0]] = append(sorted[split[0]], split[1])
		}
	}

	return sorted
}
