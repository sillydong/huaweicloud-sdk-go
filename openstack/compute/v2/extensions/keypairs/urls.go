package keypairs

import "github.com/gophercloud/gophercloud"

const resourcePath = "os-keypairs"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

// listURL generate url to list keypairs
func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

// createURL generate url to add keypair
func createURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

// getURl generate url to get keypairs
func getURL(c *gophercloud.ServiceClient, name string) string {
	return c.ServiceURL(resourcePath, name)
}

// deleteURL generate url to delete keypair
func deleteURL(c *gophercloud.ServiceClient, name string) string {
	return getURL(c, name)
}
