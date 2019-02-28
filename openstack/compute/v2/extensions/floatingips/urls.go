package floatingips

import "github.com/gophercloud/gophercloud"

const resourcePath = "os-floating-ips"

// resourceURL generate url for specified resourcePath
func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

// listURL generate url for listing servers
func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

// createURL generate url for creating server
func createURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

// getURL generate url for getting server
func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

// deleteURL generate url for deleting server
func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return getURL(c, id)
}

// serverURL generate a url for server action
func serverURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers/" + serverID + "/action")
}

// associateURL generate a url to associate floating ip with server
func associateURL(c *gophercloud.ServiceClient, serverID string) string {
	return serverURL(c, serverID)
}

// disassociateURL generate a url to disassociate floating ip with server
func disassociateURL(c *gophercloud.ServiceClient, serverID string) string {
	return serverURL(c, serverID)
}
