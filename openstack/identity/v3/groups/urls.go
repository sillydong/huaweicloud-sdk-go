package groups

import "github.com/gophercloud/gophercloud"

// listURL generate url to list groups
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

// getURL generate url to show group details
func getURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

// createURL generate url to create group
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

// updateURl generate url to update group
func updateURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

// deleteURL generate url to delete group
func deleteURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}
