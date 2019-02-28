package volumeattach

import "github.com/gophercloud/gophercloud"

const resourcePath = "os-volume_attachments"

func resourceURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, resourcePath)
}

// listURL generate URL to list attached volumes
func listURL(c *gophercloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

// createURL generate URL to attach volume
func createURL(c *gophercloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

// getURL generate URL to show attached volumes
func getURL(c *gophercloud.ServiceClient, serverID, aID string) string {
	return c.ServiceURL("servers", serverID, resourcePath, aID)
}

// deleteURL generate URL to detach volume
func deleteURL(c *gophercloud.ServiceClient, serverID, aID string) string {
	return getURL(c, serverID, aID)
}
