package flavors

import (
	"github.com/gophercloud/gophercloud"
)

// getURL generate URL to get flavor detail
func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

// detailURL generate URL for listing flavor details
func detailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

// listURL generate URL for listing flavors
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

// createURL geneate url for creating flavors
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors")
}

// deleteURL generate url for deleting flavor
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

// accessURL generate url for flavor access
func accessURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-flavor-access")
}

// accessActionURL generate url for flavor action
func accessActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "action")
}

// extraSpecsListURL generate url for flavor extra specs
func extraSpecsListURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

// extraSpecsGetURL generate url for key of flavor extra specs
func extraSpecsGetURL(client *gophercloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

// extraSpecsCreateURL generate urk for create flavor extra specs
func extraSpecsCreateURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

// extraSpecUpdateURL generate url for updating value of flavor extra specs
func extraSpecUpdateURL(client *gophercloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

// extraSpecDeleteURL generate url for delete key of flavor extra specs
func extraSpecDeleteURL(client *gophercloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}
